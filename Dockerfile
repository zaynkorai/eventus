FROM golang:1.20-bullseye

ENV DATABASE_URL=""
COPY . /app

WORKDIR /app
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /eventus /app/cmd/api/main.go

EXPOSE  8080

CMD ["/eventus"]

