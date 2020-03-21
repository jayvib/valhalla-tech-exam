GOVERSION=1.14.1
GOSOURCE=go1.14.1.linux-amd64.tar.gz

test:
	@go test -tags=unit -covermode=atomic -short ./... | grep -v '^?'

install-go:
	wget https://dl.google.com/go/${GOSOURCE}
	sha256sum ${GOSOURCE}
	tar -C /usr/local -xzf ${GOSOURCE}
	export PATH=${PATH}:/usr/local/go/bin
	go version

