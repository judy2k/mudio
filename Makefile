export GOPATH=$(abspath $(PWD)/../../../..)

all:
	go install github.com/judy2k/mudio

docker:
	GOOS=linux GOARCH=amd64 go build -o dockery/mudio github.com/judy2k/mudio
	docker build -t mudio dockery

docker-run:
	docker run -it -v $(PWD):/host mudio bash

run:
	$(GOPATH)/bin/mudio serve

clean:
	rm -f $(GOPATH)/bin/mudio
