APP_BUILD = build/bin/statApp
APP_BUILD_FAKE = build/bin/fakeApp
IMAGE = mathpoem/statApp
TAG = v1.0.1


.PHONY: build run buildFake runFake

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

push-image:
	docker push $(IMAGE):$(TAG)

run-image:
	docker run -it -p 8080:8080 -p 8888:8888 --rm $(IMAGE):$(TAG)

build-push-image: build-image push-image
