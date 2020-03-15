# Kafka with Go

Dependencies
------------

This client for Go depends on librdkafka. [Font](https://github.com/confluentinc/confluent-kafka-go).

- For MacOS X, install `librdkafka` from Homebrew
- For MacOS X, add in etc/hosts `127.0.0.1 kafka`
- For Alpine: `apk add librdkafka-dev pkgconf`

Run
---

```shell script

docker-compose up -d

cd consumer 
go run main.go

cd ..

cd producer
go run main.go

```
