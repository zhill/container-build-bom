SOURCE := cmd/bom.go
BINARY := container-bom

build: $(BINARY)

$(BINARY):
	GOOS=linux GO111MODULE=on CGO_ENABLED=0 go build -o ./bin/$(BINARY).linux_amd64 $(SOURCE)
	GOOS=darwin GO111MODULE=on CGO_ENABLED=0 go build -o ./bin/$(BINARY).darwin_amd64 $(SOURCE)

clean:
	rm ./bin/$(BINARY).*

