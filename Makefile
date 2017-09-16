export GOPATH=$(abspath $(PWD)/../../../..)

all:
	go install github.com/judy2k/mudio

run:
	$(GOPATH)/bin/mudio serve

clean:
	rm -f $(GOPATH)/bin/mudio