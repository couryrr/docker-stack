FROM golang:1.23-bookworm AS build

RUN apt-get update && \
    apt-get upgrade -y

WORKDIR /src

COPY ./ /src

RUN go build -o /bin/main .

FROM debian:bookworm-slim AS final

COPY --from=build /bin/main /bin/

CMD ["/bin/main"]

