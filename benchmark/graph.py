import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
from matplotlib import rc

rc('font', **{'family':'sans-serif','sans-serif':['DejaVu Sans'],'size':10})
rc('mathtext', **{'default':'regular'})

graph = 'w0'

plt.hlines(23, 0, 50000, linestyles='dashed', label='baseline latency', color='k')

df_op = pd.read_csv(f'set/results_op_{graph}.csv')
df_op.sort_values(by='throughput', inplace=True)
plt.plot(df_op['throughput'], df_op['latency'] / 1000000, label='op')

df_state = pd.read_csv(f'set/results_state_{graph}.csv')
df_state.sort_values(by='throughput', inplace=True)
plt.plot(df_state['throughput'], df_state['latency'] / 1000000, label='state')

df_delta = pd.read_csv(f'set/results_delta_{graph}.csv')
df_delta.sort_values(by='throughput', inplace=True)
plt.plot(df_delta['throughput'], df_delta['latency'] / 1000000, label='delta')

plt.xlabel('throughput (operations per second)')
plt.ylabel('95th percentile latency (milliseconds)')
plt.title('Throughput vs. Latency')
plt.legend()
# plt.plot(df['throughput'], df['latency'] / 1000000)
plt.savefig(f'{graph}.png', dpi=300)