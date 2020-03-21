test:
	@go test -tags=unit -covermode=atomic -short ./... | grep -v '^?'

fmt:
	go fmt ./...


