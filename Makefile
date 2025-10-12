.PHONY: fmt

fmt:
	for c in control/*.c control/*.h; do \
		clang-format --style="{BasedOnStyle: mozilla, IndentWidth: 4}" -i $$c; \
	done; \
	for go in control/*.go; do \
		go fmt $$go; \
	done; \

test:
	cd control/ && go test