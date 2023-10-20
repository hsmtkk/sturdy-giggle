FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o web

FROM gcr.io/distroless/static-debian11 AS runtime
COPY --from=builder /app/web /app/web
ENTRYPOINT ["/app/web"]
