### @@@MathPoem

Необходимо посчитать:
222
1) Сколько было потрачено gas помесячно.
2) Среднюю цену gas за день.
3) Частотное распределение цены по часам(за весь период).
4) Сколько заплатили за весь период (gas price * value).

Требования к сервису:
1) Данные должны загружаться удаленно.
2) Сервис должен вернуть все значения в виде json файла.
3) Данные должны быть посчитаны максимально быстро.

## statistic service 

    Сервис, проводящий статистический анализ на основе информации о транзакциях
    для имитации удаленного доступа к данныс реализован сервис (fakeDataSource), 
    отдающий масив нужных данных по gRPC основному сервису (statApp)
    для доступа к основному сервису реализован доступ по gRPC, а также rest api
    с помощью gRPC gateway
    
    для компиляции .proto файлов использован buf
    описание endpoint, а также описание выходных данных доступно в панели swagger
    swagger документация сгенерирована автоматически с помощью openapiv2
    
    все приложение обернуто в docker образ
    
    главная функция, вычисляющая все статистические показатели работает за время O(n)
    т.к. все значения вычисляются за один проход по массиву,
    в котором выполняется ограниченное количество действий
    

## Run

To run Docker container, execute the following:

```shell
make build-run-image
```

## Test

### gRPC

Use `grpcurl` (`curl` for gRPC) to test gRPC API:

```shell
grpcurl -plaintext localhost:50050 statApp.StatAppService/GetStat
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
