FROM node:16-buster as nodeBuilder
WORKDIR /build-staging
COPY . .
RUN make clean-full
RUN make lint-npm test-npm build-npm

FROM golang:1.16-buster as goBuilder
WORKDIR /build-staging
COPY . .
RUN make clean-full
RUN make lint-go test-go build-go

FROM debian:buster
RUN apt-get update
RUN apt-get install -y ca-certificates
WORKDIR /app
COPY --from=goBuilder /build-staging/var/personal-website ./personal-website
CMD ["./personal-website"]
EXPOSE 8000
