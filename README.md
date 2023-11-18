# go-book-store-api
Bookstore RESTful API written in Go for demo-project. Clients are GUI applications, written with C++ &amp; Java

## Software Requirements
* `Go`
* `Docker` with `docker-compose`

## Minimum Hardware Requirements
* 1 GB RAM
* 2-Core CPU
* Ethernet

## Installation
1. Build
```sh
go build
```

2. Start database
```sh
docker-compose run --build -d
```

To stop database, execute
```sh
docker-compose down
```
