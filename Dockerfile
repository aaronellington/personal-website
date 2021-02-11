FROM golang:1.15-buster as goBuilder
WORKDIR /project
COPY . .
RUN make full

FROM debian:buster
WORKDIR /project
COPY --from=goBuilder /project/var/personal-website /usr/local/bin/
COPY --from=goBuilder /project/content/ ./content
COPY --from=goBuilder /project/theme/ ./theme
CMD ["personal-website"]
EXPOSE 8080
