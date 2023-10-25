# syntax=docker/dockerfile:1

# このDockrfileを参考に、あとで作成する
# https://matsuand.github.io/docs.docker.jp.onthefly/language/golang/build-images/

##
## Build
##
FROM golang:1.21.3-bullseye AS build

WORKDIR /app

COPY ../mongoctl/go.sum ./
COPY ../mongoctl/go.mod ./
RUN go mod download

COPY ../mongoctl/importjson/*.go ./

RUN go build -o /importjson

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /importjson /importjson

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]
