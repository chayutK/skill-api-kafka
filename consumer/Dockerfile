FROM golang:alpine as builder

WORKDIR /go/src

COPY . .

RUN go mod tidy
RUN go build -o consumer

FROM alpine
COPY --from=builder /go/src/consumer .
ENTRYPOINT [ "./consumer" ]

