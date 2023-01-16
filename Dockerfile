# syntax=docker/dockerfile:1
FROM golang:1.20rc3-alpine3.17

ENV GO111MODULE=on

COPY . ./

RUN CGO_ENABLED=0 go build -o -v -a  jotone.eu

EXPOSE 8080
HEALTHCHECK --interval=5s --timeout=3s --start-period=5s --retries=3 CMD curl -f http://localhost:8080 || exit 1

CMD [ "/server" ]