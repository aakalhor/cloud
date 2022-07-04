FROM golang:1.18-alpine AS build

ARG USERNAME=TEST
ARG APIKEY=TEST

RUN cat ~/.netrc
RUN apk add git

RUN apk add librdkafka librdkafka-dev

RUN apk add gcc musl-dev

WORKDIR /tmp/go

COPY go.mod .

COPY . .

RUN go build -tags musl -o ./out/out.o .

FROM alpine:latest

COPY --from=build /tmp/go/out/out.o /app/out.o

CMD ["/app/out.o"]
