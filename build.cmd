@echo off
SET GOARCH=amd64
for %%o in (windows linux) do (
    SET GOOS=%%o
    if %%o == windows (
            go build -ldflags "-s -w" -o build/RappelzAPI.exe .\cmd\heimdall\
        ) else (  
           go build -ldflags "-s -w" -o build/RappelzAPI .\cmd\heimdall\
        )
        echo RappelzAPI-%%o built!
)