FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o muv .

FROM alpine:latest

WORKDIR /root/

RUN mkdir migration
COPY --from=builder /app/muv .
COPY --from=builder /app/migration/ migration/       

EXPOSE 8080 8080

CMD ["./muv"]