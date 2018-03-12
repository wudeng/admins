#!/bin/bash
export GOPATH=`pwd`
cd bin

if [ -f admins.exe ]; then
    rm admins.exe
fi
#d=`date -d "2010-10-18 00:00:00" +%s`
go build -o admins.exe -ldflags "-w -s -X main.version=`date +%s`" ../src/admin.go
#go build   -o admins.exe -ldflags "-w -s" ../src/admin.go
./admins.exe  -alsologtostderr -log_dir="log"


