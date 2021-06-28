import pandas as pd
import numpy as np

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
      if i * j % k == 0:
        TABLE.append('')
        continue
      else:
        TABLE.append(str(i) + 'x' + str(j) + '=' + str(i*j))
    else:
      TABLE.append(str(i) + 'x' + str(j) + '=' + str(i*j))
  times_tables.append(TABLE)

# print(idx, clm)
# print(times_tables)
df = pd.DataFrame(times_tables, columns=clm, index=idx)
df = df.T
print(df)

# nの倍数を表示しないを実装

'''sqlite3を使って表を作る'''
'''
con = sqlite3.connect('times_tables.db')
cur = con.cursor()
cur.execute("drop table if exists times_tables;")
# cur.execute("CREATE TABLE times_tables(九九 text);")
# cur.execute("INSERT INTO times_tables VALUES ?;", name)

for i in range(START, GOAL+1):
  name = str(i) + 'の段'
  if i == START:
    cur.execute("CREATE TABLE times_tables('%s')" %name)
  else:
    cur.execute("ALTER TABLE times_tables ADD COLUMN %s TEXT" %name)

''''''可変サイズのINSERTはできる？''''''
hatena = ''
for _ in range(GOAL - START):
  hatena += '?,'
hatena += '?'
for j in range(START, GOAL+1):
  cur.execute("INSERT INTO times_tables VALUES (%s);" %(hatena, times_tables[i-START][j-START]))

con.commit()
'''
'''
やり方はあってそうだけどsqliteのinsertがうまく動いていない
-> workspace/sqlite_test.pyで実験中
-> con.commit()を忘れていた
'''
