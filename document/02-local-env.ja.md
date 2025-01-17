# STEP2: 環境構築

PythonかGoから1つの言語を選び、環境を作りましょう。

---
## Pythonの環境を作る

### 1. Pythonをインストールする
* Python3.9以上がインストールされていない場合、Python3.13をインストールします
* すでに3.9以上がインストール済みの方はskipして問題ないです

### 2. Pythonのバージョンをチェックする

* インストールしたPythonにパスが通っている(=ターミナルから使える状態)か確認します。

```shell
$ python3 -V 
# もしくは $ python -V
```

表示されるPythonのバージョンがインストールしたものではなければ、**パスが通っていない**状態なので確認してください。

**:book: Reference**

* [「PATHが通っている」とはどういう状態か?](https://zenn.dev/d0ne1s/articles/1f435463551ac2)
* Windowsの方向け [【環境構築 入門】PATHの通し方 – Windows11, M1 Mac 対応 –](https://www.kikagaku.co.jp/kikagaku-blog/path/)

### 3. 依存ライブラリをインストールする

Pythonでは、`requirements.txt`というファイルに依存しているライブラリの一覧を記載します。
以下のコマンドを実行することで、依存ライブラリをまとめてインストールすることができます。

### Unix or Mac
```shell
$ cd python

# 仮想環境をつくる
$ python3 -m venv .venv

# 仮想環境をアクティベートする  
$ source .venv/bin/activate  # Unix / Macの場合
$ ./venv/Scripts/activate    # windowsの場合

# 今回使うPythonのパスを確認する。
$ which python

# 必要なライブラリをインストールする
$ pip install -r requirements.txt
```

追加でライブラリをインストールした場合は、requirements.txtにライブラリを追加するようにしましょう。

`python -m venv .venv`はPythonの仮想環境を作成するコマンドです。
仮想環境とは、プロジェクト固有のPythonの環境を作成するための方法です。
仮想環境を使うことで必要なパッケージをプロジェクトごとに分けて管理できるため、異なるプロジェクト間での依存関係の衝突を避けることができます。
仮想環境を作成したら`source .venv/bin/activate`コマンドによってその環境を有効化する必要があります。

* [venv --- 仮想環境の作成](https://docs.python.org/ja/3/library/venv.html)
* [仮想環境: Python環境構築ガイド](https://www.python.jp/install/windows/venv.html)

### 4. アプリにアクセスする

```shell
$ uvicorn main:app --reload --port 9000
```

起動に成功したら、 ブラウザで `http://127.0.0.1:9000` にアクセスして、`{"message": "Hello, world!"}`
が表示されれば成功です。

---

## Goの環境を作る
### 1. Goをインストールする
* Go1.20以上がインストールされていない場合、Go1.21をインストールします
* すでに1.20以上がインストール済みの方はskipして問題ないです

https://go.dev/dl/ このリンクからダウンロードしてください。  
※ Macの方で`x86-64`と`ARM64`どちらをダウンロードすればいいかわからない場合は、左上の🍎マーク > 「このMacについて」を開き、チップが「Apple」になっていたら`ARM64`を「Intel」であれば`x86-64`を選択してください。

### 2. Goのバージョンをチェックする

* インストールしたGoにパスが通っている(=ターミナルから使える状態)か確認します。

```shell
$ go version
```

表示されるGoのバージョンがインストールしたものではなければ、**パスが通っていない**状態なので確認してください。

**:book: Reference**

* [PATHを通すとは？- 初心者でも分かる解説](https://hara-chan.com/it/programming/environment-variable-path/)

Go関連のおすすめサイト
* [A Tour of Go](https://go.dev/tour/welcome/)
* [Go: The Complete Developer's Guide (Golang)](https://mercari.udemy.com/course/go-the-complete-developers-guide/)
  * ↑英語ですが、字幕もあり聞き取りやすいです。Section11はこのtrainingの内容と近く特に参考になると思います。


### 3. 依存ライブラリをインストールする

Goでは、`go.mod`というファイルで依存しているライブラリを管理しています。
以下のコマンドを実行することで、依存ライブラリをインストールすることができます。

```shell
$ cd go
$ go mod tidy
```

**:beginner: Point**

[このdocument](https://pkg.go.dev/cmd/go#hdr-The_go_mod_file)を参考に go.mod の役割や go.mod を扱うコマンドについて理解しましょう。

### 4. アプリにアクセスする

```shell
$ go run app/main.go
```

起動に成功したら、 ブラウザで `http://127.0.0.1:9000` にアクセスして、`{"message": "Hello, world!"}`
が表示されれば成功です。

---
**:beginner: Point**

* (LinuxやMacの場合) `.bash_profile` や `.bashrc` (zshを使っている場合は`.zshrc`)
  等はどのタイミングで呼ばれ、何をしているのか理解しましょう。
* **パスを通す** の意味を理解しましょう

**:book: Reference**

環境構築の仕方やlinuxについてさらにしっかり学ぶためには以下の教材がおすすめです。

* (JA)[book - [試して理解]Linuxのしくみ ~実験と図解で学ぶOSとハードウェアの基礎知識](https://www.amazon.co.jp/dp/477419607X/ref=cm_sw_r_tw_dp_178K0A3YTGA97XRH318R)
* (JA)[Udemy Business - もう絶対に忘れない Linux コマンド【Linux 100本ノック+名前の由来+丁寧な解説で、長期記憶に焼き付けろ！](https://mercari.udemy.com/course/linux100test/)
  * ↑わかりやすい講座だと思い貼ってますが、コマンドの暗記は特にしなくていいです

* (EN)[An Introduction to Linux Basics](https://www.digitalocean.com/community/tutorials/an-introduction-to-linux-basics)
* (EN)[Udemy Business - Linux Mastery: Master the Linux Command Line in 11.5 Hours](https://mercari.udemy.com/course/linux-mastery/)
  * You do NOT have to memorize the commands!

---
### Next

[STEP3: 出品APIを作る](03-api.ja.md)