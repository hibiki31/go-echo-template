# go-echo-template

## 作成手順

### 確認

OS確認

```
~ sw_vers
ProductName:		macOS
ProductVersion:		14.0
BuildVersion:		23A344
```

Golangのバージョン確認

```
~/g/go-echo-template go version
go version go1.21.5 darwin/amd64
```

パスを通す`vi ~/.zshrc`

```
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```


### echoのセットアップ

```bash
go mod init echo-template
# echo
go get github.com/labstack/echo/v4
# gorm
go get gorm.io/gorm
go get gorm.io/driver/postgres
# caarlos0/env 環境変数を構造体として利用する
go get github.com/caarlos0/env/v10
```

