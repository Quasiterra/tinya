TINYA_BIN = $(shell which tinya || echo /tmp/tinya)
SRC_GO := $(shell find . -name '*.go')

install: $(TINYA_BIN)

$(TINYA_BIN): $(SRC_GO)
	go install .

build: build/tinya-linux build/tinya-darwin build/tinya.exe

build/tinya-linux: $(SRC_GO)
	GOOS=linux GOARCH=amd64 go build -o $@ .

build/tinya-darwin: $(SRC_GO)
	GOOS=darwin GOARCH=amd64 go build -o $@ .

build/tinya.exe: $(SRC_GO)
	GOOS=windows GOARCH=amd64 go build -o $@ .

tests: unit-tests integration-tests

unit-tests:
	@echo "#############################"
	@echo "#    Running unit tests     #"
	@echo "#############################"
	@go test -count=1 ./core

integration-tests: install
	@echo "#############################"
	@echo "# Running integration tests #"
	@echo "#############################"
	@./test/run.sh ./test

clean:
	rm -rf build/
