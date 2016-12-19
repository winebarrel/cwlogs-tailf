PREFIX=/usr/local
VERSION=`git tag | tail -n 1`
GOOS=`go env GOOS`
GOARCH=`go env GOARCH`

ifdef GOPATH
  RUNTIME_GOPATH=$(GOPATH):`pwd`
else
  RUNTIME_GOPATH=`pwd`
endif

all: cwlogs-tailf

go-get:
	go get github.com/aws/aws-sdk-go
	go get github.com/cenkalti/backoff

cwlogs-tailf: go-get main.go src/cwlogs_tailf/optparse.go src/cwlogs_tailf/cwlogs_tailf.go
	GOPATH=$(RUNTIME_GOPATH) go build -o cwlogs-tailf main.go

install: cwlogs-tailf
	install -m 755 cwlogs-tailf $(DESTDIR)$(PREFIX)/bin/

clean:
	rm -f cwlogs-tailf *.gz

package: clean cwlogs-tailf
	gzip -c cwlogs-tailf > cwlogs-tailf-$(VERSION)-$(GOOS)-$(GOARCH).gz

deb:
	dpkg-buildpackage -us -uc
