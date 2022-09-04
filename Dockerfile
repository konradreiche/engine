FROM golang:1.18.0

RUN mkdir -p /go/src/engine

WORKDIR /go/src/engine

COPY . .

RUN go mod download

RUN go build -buildvcs=false

CMD ["./engine", "--config", "local.config.toml"]
