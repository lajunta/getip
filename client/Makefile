win:
	rsrc -manifest getip.manifest -o getip.syso
	go build -ldflags="-w -s" -o bin\getip.exe
	del getip.syso
	upx bin\getip.exe