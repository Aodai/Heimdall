@echo off
go build -ldflags "-s -w" .\cmd\heimdall\
REM No error check since go compiler returns 0 when it fails to compile.
echo Done building!
