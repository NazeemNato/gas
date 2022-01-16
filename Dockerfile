FROM go:1.17.5 as build

WORKDIR /go/src/github.com/nazeemnato/gas

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o gas main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/github.com/nazeemnato/gas .

EXPOSE 7854

CMD ["./main"]