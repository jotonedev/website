# syntax=docker/dockerfile:1

## Build
FROM golang:latest as builder

WORKDIR /

ENV GO111MODULE=on
ENV GIN_MODE=release

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -v -a -ldflags '-s -w -extldflags "-static"' -o server

## Add executable to image for healthcheck
FROM busybox:glibc as busybox

## Run
FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=busybox /bin/sh /bin/sh
COPY --from=busybox /bin/ss /bin/ss
COPY --from=busybox /bin/grep /bin/grep
COPY --from=builder /server /server

USER nonroot:nonroot
EXPOSE 8080
HEALTHCHECK --interval=5s --timeout=3s --start-period=5s --retries=3 CMD ss --listening --tcp | grep -P 'LISTEN.+:http' || exit 1

ENTRYPOINT ["/server"]
