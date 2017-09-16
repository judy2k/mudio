export GOPATH=$(abspath $(PWD)/../../../..)

all:
	go install github.com/judy2k/mudio/mudio

run:
	$(GOPATH)/bin/mudio serve