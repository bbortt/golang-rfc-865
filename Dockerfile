FROM golang:1.15.3-alpine

# https://github.com/bbortt/golang-rfc-865

CMD ["cmd"]
EXPOSE 17

WORKDIR /go/src/app

COPY . .

RUN go install -v ./...
