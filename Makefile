test:
	go test ./...

build:
	go build -o ./pass *.go

release-check:
	goreleaser check

local-release: release-check clean
	goreleaser release --snapshot --clean

ensure-no-local-changes:
	@if [ "$(shell git status -s)" != "" ]; then \
		git status -s; \
		echo "\nFATAL: refusing to release with local changes; see git status."; \
		exit 1; \
	fi

release: ensure-no-local-changes clean
	goreleaser release

clean:
	rm -rf ./pass ./dist
