package op

import (
	"fmt"
	pb "kvs/proto"
	"kvs/util"
)

type Counter struct {
	replica util.Replica
	c       int64
	current *pb.Event
}

func NewCounter(replica util.Replica) *Counter {
	return &Counter{
		replica: replica,
		c:       0,
		current: &pb.Event{
			Source:   replica.String(),
			Datatype: pb.DT_Counter,
			Data: &pb.Event_OpCounter{
				OpCounter: &pb.OpCounter{
					Update: 0,
				},
			},
		},
	}
}

func (c *Counter) Value() int {
	return int(c.c)
}

func (c *Counter) String() string {
	return fmt.Sprintf("%d", c.Value())
}

func (c *Counter) Increment() {
	c.c += 1
	update := c.current.GetOpCounter()
	update.Update += 1
}

func (c *Counter) Decrement() {
	c.c -= 1
	update := c.current.GetOpCounter()
	update.Update -= 1
}

func (c *Counter) GetEvent() *pb.Event {
	current := c.current
	c.current = &pb.Event{
		Source:   c.replica.String(),
		Datatype: pb.DT_Counter,
		Data: &pb.Event_OpCounter{
			OpCounter: &pb.OpCounter{
				Update: 0,
			},
		},
	}
	return current
}

func (c *Counter) PersistEvent(event *pb.Event) {
	oc := event.GetOpCounter()
	if oc == nil {
		fmt.Println("warning: nil opcounter encountered in PersistEvent")
		return
	}
	c.c += oc.Update
}
