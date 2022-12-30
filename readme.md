# Cat Facts

It's a microservice that collect facts about cats from [Cat Facts API](https://catfact.ninja) and store them in mongodb

## Quick Start

Run mongodb in docker

```shell
$ docker run --name mongo-storage -p 27017:27017 -d mongo
```

Build and run microservice in docker

```shell
$ docker build --tag catfacts .
$ docker run -p 3000:3000 -d catfacts
```

Or run microservice in your machine

```shell
$ make run
```
