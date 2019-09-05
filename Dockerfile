FROM golang:1.12.6-alpine3.10
RUN mkdir /app
ADD ./src /app
WORKDIR /app
RUN go build -o main