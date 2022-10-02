APP_BUILD = build/bin/statApp
APP_BUILD_FAKE = build/bin/fakeApp

.PHONY: build run buildFake runFake

build:
	clear; go build -o ${APP_BUILD} -v cmd/statApp/main.go

run: build
	${APP_BUILD}

buildFake:
	clear; go build -o ${APP_BUILD_FAKE} -v fakeDataSource/cmd/main.go

runFake: buildFake
	${APP_BUILD_FAKE}