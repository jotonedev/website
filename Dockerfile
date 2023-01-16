# syntax=docker/dockerfile:1

## Build
FROM golang:latest

ENV GO111MODULE=on

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o -v -a -ldflags '-w -extldflags "-static"' jotone.eu

## Run
FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=build /jotone.eu /server

USER nonroot:nonroot
EXPOSE 8080
HEALTHCHECK --interval=5s --timeout=3s --start-period=5s --retries=3 CMD curl -f http://localhost:8080 || exit 1

ENTRYPOINT ["/server"]
