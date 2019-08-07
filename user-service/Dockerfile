FROM golang:1.10 as builder

WORKDIR /go/src/github.com/sandeepsambidi/ekart/user-service

COPY . .

#RUN go get -u github.com/golang/dep/cmd/dep
#RUN dep init
#RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/sandeepsambidi/ekart/user-service/user-service .

CMD ["./user-service"]
