ARG GO_VERSION=1.23-alpine

FROM golang:$GO_VERSION

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr cmd/httpserver/*.go

CMD [ "./svr" ]