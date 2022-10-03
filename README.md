Необходимо посчитать:
1) Сколько было потрачено gas помесячно.
2) Среднюю цену gas за день.
3) Частотное распределение цены по часам(за весь период).
4) Сколько заплатили за весь период (gas price * value).

Требования к сервису:
1) Данные должны загружаться удаленно.
2) Сервис должен вернуть все значения в виде json файла.
3) Данные должны быть посчитаны максимально быстро.

# gRPC wrapper for rusprofile.ru

The wrapper provides access to [rusprofile.ru](https://www.rusprofile.ru/) data via gRPC. HTTP API is available via HTTP-to-gRPC gateway.

- [gRPC wrapper for rusprofile.ru](#grpc-wrapper-for-rusprofileru)
    - [Run](#run)
    - [Test](#test)
        - [gRPC](#grpc)
        - [HTTP](#http)
        - [Browser](#browser)
    - [Project structure](#project-structure)

## Run

To run Docker container, execute the following:

```shell
docker run -it -p 8080:8080 -p 8888:8888 --rm iskorotkov/rusprofile-grpc:v1.0.1
```

## Test

### gRPC

Use `grpcurl` (`curl` for gRPC) to test gRPC API:

```shell
grpcurl -plaintext -d localhost:50050 statApp.StatAppService/GetStat
```

### HTTP

Use `curl` to test HTTP API:

```shell
curl localhost:50040/stat/dayPriceAvg
```

### Browser

Open [Swagger UI](http://localhost:50040/swagger-ui/) in your browser.

## Project structure

- api - proto files and buf configuration.
    - openapiv2 - generated swagger.json file.
- build - Dockerfile.
- cmd - main package.
- pkg - generated Go files and implementation of gRPC server for getting data from rusprofile.ru.
- static/web - static files for Swagger UI page.
- tools - imports for tools used in code generation.
