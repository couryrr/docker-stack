FROM golang:1.23-bookworm

RUN apt-get update && \
    apt-get upgrade -y

WORKDIR /app

COPY ./bin/main /app

CMD ["/app/main"]

