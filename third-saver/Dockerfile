FROM golang:alpine as build

WORKDIR /src

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY third-saver.go ./

RUN go build third-saver.go

FROM alpine as runtime

COPY --from=build /src/third-saver /app/third-saver

CMD [ "/app/third-saver" ]