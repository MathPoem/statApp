APP_BUILD = build/bin/statApp

.PHONY: build run

build:
	clear; go build -o ${APP_BUILD} -v cmd/statApp/main.go

run:
	${APP_BUILD}