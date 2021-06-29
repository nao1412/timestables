import pandas as pd

print('Enter numbers...')
print('START >>', end=' ')
START = int(input())
print('GOAL >>', end=' ')
GOAL = int(input())

ESCAPE = []
print('表示しない段の有無[y/n] >>' ,end =' ')
judge = input()
if judge == 'y':
  print('ESCAPE (separated by spaces) >>', end=' ')
  ESCAPE = list(map(int, input().split()))

nnn = []
mmm = []
print('積がnの倍数のみ表示[y/n], nの倍数を表示しない場合[x] (separated by spaces) >>', end=' ')
judge = input()
if judge == 'y':
  print('n (separated by spaces) >>', end=' ')
  nnn = list(map(int, input().split()))
elif judge == 'x':
  print('n (separated by spaces) >>', end=' ')
  mmm = list(map(int, input().split()))

times_tables = []
clm = []; idx = []
for i in range(START, GOAL+1):
  clm.append('x' + str(i))
  if i in ESCAPE: continue
  TABLE = []
  idx.append(str(i) + ' times')
  for j in range(START, GOAL+1):
    if len(nnn) > 0:
      for k in nnn:
        if i * j % k == 0:
          TABLE.append(str(i) + 'x' + str(j) + '=' + str(i*j))
        else:
          TABLE.append('')
    elif len(mmm) > 0:
      for k in mmm:
        if i * j % k == 0:
          TABLE.append('')
        else:
          TABLE.append(str(i) + 'x' + str(j) + '=' + str(i*j))
    else:
      TABLE.append(str(i) + 'x' + str(j) + '=' + str(i*j))
  times_tables.append(TABLE)

df = pd.DataFrame(times_tables, columns=clm, index=idx)
df = df.T
print(df)