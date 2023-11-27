BINARY = service
SERVICE = family_service
PORT = 80

build:
	CGO_ENABLED=0 GOOS=linux go build -o $(BINARY) ./src/main.go 

run:
	./$(BINARY)

clean:
	rm -f $(BINARY)



docker-build:
	docker build -t $(SERVICE) .

docker-run:
	docker run -p $(PORT):$(PORT) $(SERVICE)

docker-push:
	docker push $(SERVICE)
