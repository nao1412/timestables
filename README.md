# Times tables & Notice board

値の範囲を変えられる九九表(problem1)と掲示板(problem2)のコードです
 
# DEMO
 
![image](https://user-images.githubusercontent.com/64777602/124921100-05c9d780-e033-11eb-923e-a160ba1242dd.png)

 problem1では`index.html`をブラウザで開くことで上のような九九表を得られます。
 
 
# Features
 
* problem1

`START`(default 1)と`GOAL`(default 9)に値(半角)を入力し`Run`をクリックことで、(GOAL-START+1)x(GOAL-START+1)の掛け算の表を作ることができます。範囲はありませんが現実的な計算量では200x200程度なら10秒ほどで出力されます。

また、`Delete table`(default 0)に値を入力すれば、その段を出力結果での表示をなくすことができます。

`Only multiples of N`(default 1)は入力値Nの倍数のみを出力できるようになり、`Except multiples of N`(default -1)では入力値Nの倍数を出力結果で空白にすることができます。


※注意

`Run`をリロードせずに二回連続で押すと表をうまく出力できません。値を変更したい場合はブラウザをリロードするか、`Reset`ボタンをクリックすることで初期化してください。

それぞれの値は複数入力することはできません。1つずづ入力してください。

* problem2
 
# Requirement
 
* Go
 
# Installation
 
 
```bash
```
 
# Usage
 
```
 
# Note
 
注意点などがあれば書く

# References
 
# Author
 
* 二宮直也
 
