
arm:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-extldflags=-static" -o bin/arm/devmemgo

arm64:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-extldflags=-static" -o bin/arm64/devmemgo

amd64:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-extldflags=-static" -o bin/amd64/devmemgo

clean:
	rm bin/arm64/devmemgo bin/arm/devmemgo bin/amd64/devmemgo

.PHONY: run
run:
	go run main.go