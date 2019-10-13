FROM golang:1.13-alpine3.10 AS builder

ENV GOPROXY=https://proxy.golang.org

WORKDIR /project

COPY go.mod .
COPY go.sum .
RUN go mod download

# tzdata is required for clickhouse client
RUN apk add --update --no-cache gcc libc-dev tzdata

COPY cmd/server ./cmd/server
COPY pkg ./pkg

RUN go build -tags clz4 -ldflags="-s -w -extldflags '-static'" -o /report-aggregator ./cmd/server

FROM scratch

ENV SERVER_PORT=80

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /report-aggregator .
EXPOSE 80
ENTRYPOINT ["/report-aggregator", "--db", "clickhouse:9000"]