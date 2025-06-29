all: build test dist

build:
	@mkdir -p artifacts
	go build -o artifacts/xgoimports .

test: build
	cd tests && make

install:
	go install -v

uninstall:
	@LOCAL_BINARY="$$(go env GOPATH)/bin/xgoimports"; \
	if [ -f "$$LOCAL_BINARY" ]; then \
		rm -f "$$LOCAL_BINARY"; \
		echo "Uninstalled $$LOCAL_BINARY"; \
	else \
		echo "No binary found at $$LOCAL_BINARY"; \
	fi

dist: build
	@mkdir -p artifacts/dist
	cp artifacts/xgoimports artifacts/dist/xgoimports
	cp LICENSE artifacts/dist/LICENSE
	cp NOTICE artifacts/dist/NOTICE
	cp README.md artifacts/dist/README.md

release:
	goreleaser release --snapshot --clean
