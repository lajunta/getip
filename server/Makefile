grpc:
	protoc --go_out=. --go-grpc_out=.  grpcd/getip.proto

win:
	go build -ldflags="-w -s" -o bin/getipd.exe
	upx bin/getipd.exe