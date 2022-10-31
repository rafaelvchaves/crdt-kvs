package worker

import (
	"context"
	"kvs/crdt"
	"kvs/crdt/generator"
	pb "kvs/proto"
	"kvs/util"
	"os"
	"time"

	"github.com/buraksezer/consistent"
	"google.golang.org/grpc"
)

const RPCTimeout = 10 * time.Second

type CRDTHandler interface {
	Broadcast(*pb.Event)
}

type Worker[F crdt.Flavor] struct {
	replica     util.Replica
	generator   generator.Generator[F]
	kvs         Store[F]
	requests    chan ClientRequest
	events      chan *pb.Event
	hashRing    *consistent.Consistent
	connections map[string]*grpc.ClientConn
	cfg         util.Config
	logger      *util.Logger
}

type Operation int

const (
	Increment Operation = iota
	Decrement
	Get
)

type ClientRequest struct {
	Key       string
	Operation Operation
	Params    []string
	Response  chan Response
}

type Response = struct {
	Value  string
	Exists bool
}

func New[F crdt.Flavor](replica util.Replica, generator generator.Generator[F], logger *util.Logger) *Worker[F] {
	return &Worker[F]{
		generator:   generator,
		replica:     replica,
		kvs:         NewCache[F](),
		requests:    make(chan ClientRequest, 10),
		events:      make(chan *pb.Event, 10),
		hashRing:    util.GetHashRing(),
		connections: util.GetConnections(),
		cfg:         util.LoadConfig(),
		logger:      logger,
	}
}

func (w *Worker[F]) Get(key string) (string, bool) {
	v, ok := w.kvs.Get(key)
	if ok {
		return v.String(), true
	}
	return "", false
}

func (w *Worker[F]) Start() {
	requestDeadline := 50 * time.Millisecond
	for {
		// set of keys modified in this epoch
		changeset := util.NewSet[string]()
	reqLoop:
		for {
			// phase 1: receive client requests and convert to events
			select {
			case req := <-w.requests:
				w.process(req)
				if req.Operation != Get {
					changeset.Add(req.Key)
				}
			case <-time.After(requestDeadline):
				break reqLoop
			}
		}
		// phase 2: go through all affected keys and broadcast to other owners
		changeset.Range(func(key string) bool {
			v, ok := w.kvs.Get(key)
			if !ok {
				return true
			}
			e := v.GetEvent()
			e.Key = key
			w.broadcast(e)
			return true
		})

	eventLoop:
		for {
			select {
			case event := <-w.events:
				w.logger.Infof("worker %d received (%q, %+v) from %q\n", w.replica.WorkerID, event.Key, event.Data, event.Source)
				v := w.kvs.GetOrDefault(event.Key, w.generator.New(event.Datatype, w.replica))
				v.PersistEvent(event)
			case <-time.After(requestDeadline):
				break eventLoop
			}
		}
	}
}

func (w *Worker[F]) process(r ClientRequest) {
	switch r.Operation {
	case Increment:
		w.logger.Infof("replica %s handling increment on key %q", w.replica.String(), r.Key)
		v := w.kvs.GetOrDefault(r.Key, w.generator.New(pb.DT_Counter, w.replica))
		counter, ok := v.(crdt.Counter)
		if !ok {
			return
		}
		counter.Increment()
		w.kvs.Put(r.Key, v)
	case Decrement:
		w.logger.Infof("replica %s handling decrement on key %q", w.replica.String(), r.Key)
		v := w.kvs.GetOrDefault(r.Key, w.generator.New(pb.DT_Counter, w.replica))
		counter, ok := v.(crdt.Counter)
		if !ok {
			return
		}
		counter.Decrement()
		w.kvs.Put(r.Key, v)
	case Get:
		w.logger.Infof("replica %s handling get on key %q", w.replica.String(), r.Key)
		var value string
		v, ok := w.kvs.Get(r.Key)
		if ok {
			value = v.String()
		}
		r.Response <- Response{
			Value:  value,
			Exists: ok,
		}
	}
	// fmt.Printf("%s: %s\n", w.replica.String(), w.kvs.String())
}

func (w *Worker[F]) broadcast(event *pb.Event) {
	owners, err := w.hashRing.GetClosestN([]byte(event.Key), w.cfg.RepFactor)
	if err != nil {
		os.Exit(1)
	}
	for _, o := range owners {
		v := o.(util.Replica)
		if v == w.replica {
			continue
		}
		event.Dest = int32(v.WorkerID)
		client := pb.NewChiaveClient(w.connections[v.Addr])
		ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
		defer cancel()
		w.logger.Infof("worker %d is sending %v to worker %d", w.replica.WorkerID, event.Data, v.WorkerID)
		_, err := client.ProcessEvent(ctx, event)
		if err != nil {
			w.logger.Errorf("ProcessEvent from %s to %s: %v", w.replica.String(), v.String(), err)
		}
	}
}

func (w *Worker[_]) PutRequest(r ClientRequest) {
	w.requests <- r
}

func (w *Worker[F]) PutEvent(e *pb.Event) {
	w.events <- e
}
