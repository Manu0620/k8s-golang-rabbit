FROM golang:alpine as build

WORKDIR /src

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY second-consumer.go ./

RUN go build second-consumer.go

FROM alpine as runtime

COPY --from=build /src/second-consumer /app/second-consumer

CMD [ "/app/second-consumer" ]