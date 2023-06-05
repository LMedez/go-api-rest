VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "1.0.0")
LDFLAGS := -ldflags "-X main.Version=${VERSION}"

PID_FILE := './.pid'

.PHONY: run-restart
run-restart: ## restart the API server
	@pkill -P `cat $(PID_FILE)` || true
	@printf '%*s\n' "80" '' | tr ' ' -
	@echo "Source file changed. Restarting server..."
	@go run ${LDFLAGS} cmd/main.go & echo $$! > $(PID_FILE)
	@printf '%*s\n' "80" '' | tr ' ' -

run-live: ## run the API server with live reload support (requires fswatch)
	@go run ${LDFLAGS} cmd/main.go & echo $$! > $(PID_FILE)
	@fswatch -x -o --event Created --event Updated --event Renamed -r internal pkg cmd config | xargs -n1 -I {} make run-restart