FROM golang:alpine AS builder

ADD . /src
RUN cd /src && go build -o main.bin cmd/main.go

FROM alpine as runner

WORKDIR /app
COPY --from=builder /src/main.bin /app/
COPY .env /app/.env
COPY assets /app/assets

ENTRYPOINT ./main.bin

CMD ["/app"]