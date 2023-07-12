FROM golang:1.20-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o app .

FROM rust:latest
RUN rustup target add wasm32-unknown-unknown
RUN rustup component add rustfmt
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 3001
CMD ["./app"]
