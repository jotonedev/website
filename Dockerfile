# syntax=docker/dockerfile:1

## Build
FROM golang:latest as builder

WORKDIR /build
ENV GO111MODULE=on
ENV GIN_MODE=release

COPY go.* ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -v -a -ldflags '-s -w -extldflags "-static"' -o server


## Run
FROM busybox:latest

WORKDIR /srv
ENV GIN_MODE=release

RUN echo "nonroot:x:1002:1002:nobody:/:/bin/sh" >> /etc/passwd
RUN echo "nonroot:x:1002:" >> /etc/group
COPY --from=builder /build/server /srv/server
RUN chmod 005 /srv/server

USER nonroot:nonroot
EXPOSE 8080
HEALTHCHECK --interval=5s --timeout=3s --start-period=5s --retries=3 CMD /bin/netstat -atl | /bin/grep ':8080' || exit 1

CMD ["/srv/server"]
