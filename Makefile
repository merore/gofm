BIN=bin
CMD=./cmd
IMAGE=merore/gofm
VERSION?=latest

LDFLAGS="-s -w -extldflags '-static'"

.PHONY: all
all: build

.PHONY: build
build:
	@for binary in $(CMD)/*; do \
		echo "Building $$(basename $$binary)"; \
		CGO_ENABLED=0 go build -v -o $(BIN)/$$(basename $$binary) -ldflags $(LDFLAGS) $$binary;\
	done

.PHONY: image
image: build
	@docker build -t $(IMAGE):$(VERSION) .

.PHONY: clean
clean:
	@rm -rf $(BIN)

.PHONY: env
env:
	./bin/gofm --live="$(MISSEVAN_LIVE)" --missevan-token="$(MISSEVAN_TOKEN)" --openai-token="$(OPENAI_TOKEN)"