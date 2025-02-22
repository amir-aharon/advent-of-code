
ROOT_DIR := $(shell git rev-parse --show-toplevel 2>/dev/null || realpath $(dir $(lastword $(MAKEFILE_LIST))))

default: help

test:
	@if [ -z "$(DAY)" ]; then echo "Error: DAY is required. Usage: make test DAY=<day_number>"; exit 1; fi
	@$(MAKE) -C $(ROOT_DIR) _test_internal DAY=$(DAY)

_test_internal:
	@cd 2024/day$(DAY) && sed -i.bak 's/"input.txt"/"test.txt"/g' main.go && go run main.go && mv main.go.bak main.go

run:
	@if [ -z "$(DAY)" ]; then echo "Error: DAY is required. Usage: make run DAY=<day_number>"; exit 1; fi
	@$(MAKE) -C $(ROOT_DIR) _run_internal DAY=$(DAY)

_run_internal:
	@cd 2024/day$(DAY) && go run main.go

help:
	@echo "Usage:"
	@echo "  make test DAY=<day>  # Runs 'go run main.go' with 'input.txt' replaced by 'test.txt'"
	@echo "  make run DAY=<day>   # Runs 'go run main.go' normally"
	@echo "  make                 # Prints this help message"
