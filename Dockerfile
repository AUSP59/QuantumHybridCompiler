FROM golang:1.22
WORKDIR /app
COPY . .
RUN go build -o qhc ./cmd/qhc
CMD ["./qhc"]
