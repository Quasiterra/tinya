TINYA_BIN = $(shell which tinya || echo /tmp/tinya)

install: $(TINYA_BIN)

$(TINYA_BIN): $(shell find . -name '*.go')
	go install .

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
