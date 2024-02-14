clean:
	rm bin/server
	rm bin/client

build-server: clean
	go build -o bin/server .

build-client: clean
	go build -o bin/client client/main.go

run-server-dev:
	go run main.go

run-client-dev:
	go run client/main.go

run-server:
	./bin/server

run-client:
	./bin/client

test:
	go test -v -cover test/model/queue_linked_list_test.go  
