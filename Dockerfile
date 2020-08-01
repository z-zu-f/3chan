FROM golang:alpine AS go-builder
WORKDIR /go/src/3chan
ADD . /go/src/3chan
RUN go build -o ./main;

FROM alpine:latest
WORKDIR /go/src/3chan
COPY --from=go-builder /go/src/3chan/views ./views
COPY --from=go-builder /go/src/3chan/main .
CMD ["/go/src/3chan/main"]