APP_BUILD = build/bin/statApp
APP_BUILD_FAKE = build/bin/fakeApp
IMAGE = stat-app
TAG = v1


.PHONY: build run buildFake runFake build-image run-image build-push-image

build:
	clear; go build -o ${APP_BUILD} -v cmd/statApp/main.go

run: build
	${APP_BUILD}

buildFake:
	clear; go build -o ${APP_BUILD_FAKE} -v fakeDataSource/cmd/main.go

runFake: buildFake
	${APP_BUILD_FAKE}

build-image:
	docker build -t $(IMAGE):$(TAG) -f ./build/statApp.dockerfile .

run-image:
	docker run -it -p 50040:50040 -p 50050:50050 --rm $(IMAGE):$(TAG)

build-run-image: build-image run-image
