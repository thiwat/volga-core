FROM golang:1.18.10-alpine3.17 AS build_base

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/bin .


FROM alpine:3.17
RUN apk add ca-certificates

COPY --from=build_base /app/out/bin /app/bin
COPY .env .

EXPOSE 3000

CMD ["/app/bin"]