FROM golang:1.20-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM rust:latest
RUN rustup target add wasm32-unknown-unknown
RUN rustup component add rustfmt
RUN curl -LJO https://github.com/mozilla/sccache/releases/download/v0.4.1/sccache-v0.4.1-x86_64-unknown-linux-musl.tar.gz
RUN tar -xvf sccache-v0.4.1-x86_64-unknown-linux-musl.tar.gz
RUN mv sccache-v0.4.1-x86_64-unknown-linux-musl/sccache /usr/local/bin
ENV RUSTC_WRAPPER=/usr/local/bin/sccache
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 3001
CMD ["./app"]
