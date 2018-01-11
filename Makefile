SHELL          := /bin/bash
VERSION        := v0.1.4
GOOS           := $(shell go env GOOS)
GOARCH         := $(shell go env GOARCH)
RUNTIME_GOPATH := $(GOPATH):$(shell pwd)
SRC            := $(wildcard *.go) $(wildcard src/*/*.go)

CENTOS_IMAGE=docker-go-pkg-build-centos6
CENTOS_CONTAINER_NAME=docker-go-pkg-build-centos6-$(shell date +%s)

all: cwlogs-tailf

cwlogs-tailf: go-get $(SRC)
	GOPATH=$(RUNTIME_GOPATH) go build -a -tags netgo -installsuffix netgo -o cwlogs-tailf
ifeq ($(GOOS),linux)
	[[ "`ldd cwlogs-tailf`" =~ "not a dynamic executable" ]] || exit 1
endif

go-get:
	go get github.com/aws/aws-sdk-go
	go get github.com/cenkalti/backoff
	go get github.com/tkuchiki/parsetime

package: clean cwlogs-tailf
	gzip -c cwlogs-tailf > pkg/cwlogs-tailf-$(VERSION)-$(GOOS)-$(GOARCH).gz

package\:linux:
	docker run --name $(CENTOS_CONTAINER_NAME) -v $(shell pwd):/tmp/src $(CENTOS_IMAGE) make -C /tmp/src package
	docker rm $(CENTOS_CONTAINER_NAME)

rpm:
	docker run --name $(CENTOS_CONTAINER_NAME) -v $(shell pwd):/tmp/src $(CENTOS_IMAGE) make -C /tmp/src rpm:docker
	docker rm $(CENTOS_CONTAINER_NAME)

rpm\:docker: clean
	cd ../ && tar zcf cwlogs-tailf.tar.gz src
	mv ../cwlogs-tailf.tar.gz /root/rpmbuild/SOURCES/
	cp cwlogs-tailf.spec /root/rpmbuild/SPECS/
	rpmbuild -ba /root/rpmbuild/SPECS/cwlogs-tailf.spec
	mv /root/rpmbuild/RPMS/x86_64/cwlogs-tailf-*.rpm pkg/
	mv /root/rpmbuild/SRPMS/cwlogs-tailf-*.src.rpm pkg/

docker\:build\:centos6:
	docker build -f docker/Dockerfile.centos6 -t $(CENTOS_IMAGE) .

clean:
	rm -f pkg/*
