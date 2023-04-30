FROM golang:1.20-alpine as base
WORKDIR /app
COPY . .

FROM base as test
RUN go test /app/...

FROM base as build
RUN go build /app/cmd/urlshorten

FROM alpine:latest
WORKDIR /app
EXPOSE 10000
COPY --from=build /app/urlshorten .
ENTRYPOINT ["/app/url_shorten_back"]
