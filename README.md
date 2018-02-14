# tester
[![Build Status](https://travis-ci.org/frankgreco/tester.svg?branch=master)](https://travis-ci.org/frankgreco/tester) [![Docker Repository on Quay](https://quay.io/repository/frankgreco/tester/status "Docker Repository on Quay")](https://quay.io/repository/frankgreco/tester)

> a tester go application with prometheus instrumentation.

## Quick Start
```sh
$ mkdir -p $GOPATH/src/frankgreco
$ cd $GOPATH/src/frankgreco
$ git clone git@github.com:frankgreco/tester.git
$ cd tester
$ make
$ ./tester &
$ curl http://localhost:8080
```
## Docker
```sh
docker run -d -p 8080:8080 -p 9000:9000 quay.io/frankgreco/tester
$ curl http://localhost:8080
```

## Prometheus
Multiple metrics are exposed by this tester application. These metrics can be scraped by Prometheus at the following endpoint.
```
$ curl http://localhost:9000
```