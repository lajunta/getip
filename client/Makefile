win:
	rsrc -manifest getip.manifest -o getip.syso
	go build -ldflags="-w -s" -o getip.exe
	del getip.syso
	upx getip.exe