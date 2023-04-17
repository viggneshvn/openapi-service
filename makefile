BINARY_NAME=openapi-service

build:
	go build -o $(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

.PHONY: build clean run

generate:
	openapi-generator generate -i openapi.yaml -g go -o ./gen