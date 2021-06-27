import sqlite3

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
print('積がnの倍数のみ表示[y/n] >>', end=' ')
judge = input()
if judge == 'y':
  print('n (separated by spaces) >>', end=' ')
  nnn = list(map(int, input().split()))

times_tables = []
for i in range(START, GOAL+1):
  if i in ESCAPE: continue
  DAN = []
  for j in range(START, GOAL+1):
    if len(nnn) > 0:
      for k in nnn:
        if i * j % k == 0:
          print('%d×%d=%d' %(i, j , i*j), end=' ')
          DAN.append(str(i) + '×' + str(j) + '=' + str(i*j))
        else: 
          print('     ', end=' ')
          DAN.append('     ')
    else:
      print('%d×%d=%d' %(i, j , i*j), end=' ')
      DAN.append(str(i) + '×' + str(j) + '=' + str(i*j))
    if j == GOAL: print('')
  times_tables.append(DAN)

print(times_tables)
name = times_tables[0][0]
print(name)


'''sqlite3を使って表を作る'''
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

'''可変サイズのINSERTはできる？'''
hatena = ''
for _ in range(GOAL - START):
  hatena += '?,'
hatena += '?'
for j in range(START, GOAL+1):
  cur.execute("INSERT INTO times_tables VALUES (%s);" %(hatena, times_tables[i-START][j-START]))

con.commit()

'''
やり方はあってそうだけどsqliteのinsertがうまく動いていない 
-> workspace/sqlite_test.pyで実験中
-> con.commit()を忘れていた
'''
