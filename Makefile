test:
	@go test -tags=unit -covermode=atomic -short ./... | grep -v '^?'

fmt:
	go fmt ./...

build:
	mkdir -p ./bin/
	for i in $$(ls); do \
		if [ -d $${i} ]; then \
			if [ $${i} != "bin" ]; then \
				echo "Building $${i} -> ./bin/$${i}.exe"; \
				go build -o ./bin/$${i}.exe ./$${i}; \
			fi; \
		fi; \
	done




