all: linter mod build test

help: build # Show help
	./kafka-client --help

build: # build code
	go build -o kafka-client main.go

docker: linter # Create a docker image
	echo "Create the kafka-client image"
	docker build -t kafka-client .

mod: # update dependencies before commit
	go mod tidy
	go mod vendor

test: test-unit test-version test-help

test-unit: # go test unit cases
	go test -v ./...

test-sendMsg: build # test sending 1M messages
	./kafka-client sendMsg \
        --kafka-url "kafka0:9092" \
        --topic "test" \
        --num-msg 1000000 \
        --batch-size 1000

test-listTopics: build # test listing existing topics
	./kafka-client listTopics \
        --kafka-url "kafka0:9092"

test-getOffset: build # test getting the last offset
	./kafka-client getOffset \
        --kafka-url "kafka0:9092" \
        --topic "test"

test-version: # test showing version number
	go run -race main.go version

test-help: # test showing help information
	./kafka-client help

linter: # format code before commit
	for d in cmd kafka; do \
		gofmt -l -s -w $$d/*.go ;\
		goimports -l -w $$d/*.go ;\
		go vet $$d/*.go ;\
		golint $$d/*.go ;\
	done

clean: # clean built binaries
	rm -f ./kafka-client
	rm -f ${GOPATH}/bin/kafka-client

.PHONY: \
	build \
	clean \
	docker \
	mod \
	linter
