FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .
RUN go test /app/...
RUN go build /app


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/url_shorten_back .
CMD ["/app/url_shorten_back"]
