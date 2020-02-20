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