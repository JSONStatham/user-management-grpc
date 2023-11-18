.PHONY: proto
proto: clean format gen lint

BUF_VERSION=1.28.0

.PHONY: gen
gen:
	@$(GOPATH)/bin/buf generate
	@cd $(CURDIR)/gen/go && \
	go mod init github.com/JSONStatham/user-management-grpc/gen/go && go mod tidy; \

.PHONY: lint
lint:
	@$(GOPATH)/bin/buf lint

.PHONY: format
format:
	@$(GOPATH)/bin/buf format

.PHONY: clean
clean:
	@rm -rf ./gen || true
