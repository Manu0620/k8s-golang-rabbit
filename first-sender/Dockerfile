FROM golang:alpine as build

WORKDIR /src

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY first-sender.go ./

RUN go build first-sender.go

FROM alpine as runtime

COPY --from=build /src/first-sender /app/first-sender 

CMD [ "/app/first-sender" ]