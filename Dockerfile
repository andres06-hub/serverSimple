FROM golang:1.20-alpine
WORKDIR /app
COPY ./go.* .
RUN go mod download
COPY . .
RUN go build -o main .
CMD ["./main"]