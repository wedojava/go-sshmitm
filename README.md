# go-sshmitm



build command:

- xgo

```powershell
go get github.com/karalabe/xgo
xgo -out dist/go-sshmitm ./
```

- go built-in

```powershell
set GOARCH=amd64
set GOOS=windows
go build -o dist/go-sshmitm-windows-amd64.exe ./

set GOARCH=386
set GOOS=windows
go build -o dist/go-sshmitm-windows-386.exe ./

set GOARCH=amd64
set GOOS=linux
go build -o dist/go-sshmitm-linux-amd64 ./

set GOARCH=386
set GOOS=linux
go build -o dist/go-sshmitm-linux-386 ./

set GOARCH=arm64
set GOOS=linux
go build -o dist/go-sshmitm-linux-arm64 ./

set GOARCH=arm
set GOOS=linux
go build -o dist/go-sshmitm-linux-arm ./

set GOARCH=arm64
set GOOS=linux
go build -o dist/go-sshmitm-linux-arm64 ./

set GOARCH=arm
set GOOS=linux
go build -o dist/go-sshmitm-linux-arm ./
```



Reference: [Golang交叉编译各个平台的二进制文件](https://studygolang.com/articles/14376)

GOOS和GOARCH支持列表

| GOOS - Target Operating System | GOARCH - Target Platform |
| ------------------------------ | ------------------------ |
| android                        | arm                      |
| darwin                         | 386                      |
| darwin                         | amd64                    |
| darwin                         | arm                      |
| darwin                         | arm64                    |
| dragonfly                      | amd64                    |
| freebsd                        | 386                      |
| freebsd                        | amd64                    |
| freebsd                        | arm                      |
| linux                          | 386                      |
| linux                          | amd64                    |
| linux                          | arm                      |
| linux                          | arm64                    |
| linux                          | ppc64                    |
| linux                          | ppc64le                  |
| linux                          | mips                     |
| linux                          | mipsle                   |
| linux                          | mips64                   |
| linux                          | mips64le                 |
| netbsd                         | 386                      |
| netbsd                         | amd64                    |
| netbsd                         | arm                      |
| openbsd                        | 386                      |
| openbsd                        | amd64                    |
| openbsd                        | arm                      |
| plan9                          | 386                      |
| plan9                          | amd64                    |
| solaris                        | amd64                    |
| windows                        | 386                      |
| windows                        | amd64                    |

