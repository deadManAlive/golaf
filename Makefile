SOURCES := $(wildcard main.go util/*.go)
TESTFILE := sine.wav

.PHONY: build check

build: $(SOURCES)
	go build

check: build
	./golaf.exe $(TESTFILE)