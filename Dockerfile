FROM golang:1.20-alpine as base
WORKDIR /app
COPY . .

FROM base as test
RUN go test /app/...

FROM base as build
RUN go build /app

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/url_shorten_back .
CMD ["/app/url_shorten_back"]
