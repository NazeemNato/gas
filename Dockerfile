# Building the binary of the App
FROM golang:1.17 AS build

WORKDIR /go/src/gas

# Copy all the Code and stuff to compile everything
COPY . .


RUN go mod download
ENV CGO_ENABLED=0
ENV GOOS=linux 
ENV GOARCH=amd64
RUN go build -o main .


FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/gas/main .

# Exposes port 3000 because our program listens on that port
EXPOSE 3000

CMD ["./main"]