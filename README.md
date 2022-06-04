# practice_go
## 学習内容に関するまとめ
### _main.go
・goパッケージに関する知識<br>
・context<br>
・crypto/hmac<br>
・encoding/json<br>
・io/ioutil<br>
・net/http<br>
・net/url<br>
・regexp<br>
・sort<br>
・iotaの使い方<br>
・structのjson設定方法（ex omitempty...<br>
・パッケージの関数をoverride<br>

### __main.go
・golang.org/x/sync/semaphore<br>

### ___main.go
・ini<br>
・bitflyerからBTCの値段を取得<br>

### ____main.go
・DB操作<br>


## コマンド関連
gofmt使用
```
$ gofmt example.go
```
実際にファイルをフォーマットしたもので上書き
```
$ gofmt -w example.go
```

go test使用
```
$ go test ./...
```

godoc確認
```
$ go doc fmt Println
```

sqlite3の操作
起動
```
$ sqlite3
```
終了
```
sqlite3> .exit
```
