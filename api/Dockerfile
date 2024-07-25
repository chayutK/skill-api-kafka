FROM golang:alpine AS builder

WORKDIR /go/src

COPY . .

RUN go mod tidy
RUN go build -o app

FROM alpine
COPY --from=builder /go/src/app .
ENTRYPOINT [ "./app" ]

