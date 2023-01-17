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


## Run
FROM gcr.io/distroless/static-debian11:debug

WORKDIR /

COPY --from=builder /server /server

USER nonroot:nonroot
EXPOSE 8080
HEALTHCHECK --interval=5s --timeout=3s --start-period=5s --retries=3 CMD /bin/netstat -atl | /bin/grep -P '0.0.0.0:http' || /bin/exit 1

ENTRYPOINT ["/server"]
