build:
	docker build -t family-team .

run:
	docker run -p 80:80 family-team

test:
	go test -v ./...
