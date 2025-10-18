all: test

fmt:
	for c in control/*.c control/*.h; do \
		clang-format --style="{BasedOnStyle: mozilla, IndentWidth: 4}" -i $$c; \
	done; \
	go fmt ./...

test:
	cd test/ && go test

example:
	go run ./examples/control/

.PHONY: all fmt test