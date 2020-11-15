# microservice
Setup for a HTTP service

## features

- Simple HTTP API, basic handler tests
- Httprouter
- Makefile
- multi-stage builds: see Dockerfile
- Structured logging: using logrus
- Versioning: service has version, commit and build time
- Health checks

## usage

Build and run a docker image:
```sh
make run
```
Make a request to the service:
```sh
curl -i localhost:8000/
HTTP/1.1 200 OK
Date: Sun, 15 Nov 2020 22:10:10 GMT
Content-Length: 83
Content-Type: text/plain; charset=utf-8

Request path:/
```

### Health checks:
1. `/health/live`
2. `/health/ready`
