SOURCES := $(wildcard main.go util/*.go)
TESTFILE := sine.wav

build: $(SOURCES)
	go build

check: build
	./golaf.exe $(TESTFILE)