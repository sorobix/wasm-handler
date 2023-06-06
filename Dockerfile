FROM golang:1.18-alpine

WORKDIR /app

RUN apk update && apk add --no-cache curl gcc musl-dev rust cargo rustfmt

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

EXPOSE 3001

CMD ["./app"]