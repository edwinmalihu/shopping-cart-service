FROM golang:1.19-alpine3.15 AS build_base

WORKDIR /tmp/api
RUN apk add git
COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:3.9 
RUN apk update && apk add tzdata
WORKDIR /app

RUN chgrp -R 0 /app && \
    chmod -R g=u /app

ENV TZ=Asia/Jakarta
COPY --from=build_base /tmp/api/main api
COPY --from=build_base /tmp/api/config config

CMD ["./api"]