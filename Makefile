.PHONY: buildUbuntu
buildUbuntu:
	go build -o interview cmd/main.go

.PHONY: buildWindows
buildWindows:
	go build -o interview.exe cmd/main.go

.PHONY: runUbuntu
runUbuntu:
	go build -o interview cmd/main.go
	./interview

.PHONY: runWindows
runWindows:
	go build -o interview.exe cmd/main.go
	.\interview.exe

.PHONY: test
test:
	go test test/db_test.go

DEFAULT_GOAL := runUbuntu