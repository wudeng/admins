.PHONY: all run

IDENTITY='c:/windows/id_rsa'
HOST=ts@192.168.4.168
ADMINS='/data/server/admins/'



all:
	$(shell \
		export GOPATH=`pwd`; \
		export GOARCH=amd64; \
		export GOOS=linux; \
		cd bin; \
		go build -o admins -ldflags "-w -s -X main.version=`date +%s`" ../src/admin.go)

run:
	tar zcf admin.tar.gz ./bin
	scp -i $(IDENTITY) admin.tar.gz $(HOST):$(ADMINS)
	ssh -i $(IDENTITY) $(HOST) "cd $(ADMINS); tar zxf admin.tar.gz; cd bin; chmod +x admins; pkill admins; bash run"
