print('Enter numbers...')
print('START >>', end=' ')
START = int(input())
print('GAOL >>', end=' ')
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

for i in range(START, GOAL+1):
  if i in ESCAPE: continue
  for j in range(START, GOAL+1):
    if len(nnn) > 0:
      for k in nnn:
        if i * j % k == 0:
          print('%d×%d=%d' %(i, j , i*j), end=' ')
        else: print('     ', end=' ')
    else:
      print('%d×%d=%d' %(i, j , i*j), end=' ')
    if j == GOAL: print('')

'''sqliteを使って表を作る'''