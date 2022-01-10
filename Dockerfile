FROM golang:latest
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build ./cmd/band-archive-api-server
EXPOSE 8080
CMD ["./band-archive-api-server"]