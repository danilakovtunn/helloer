FROM golang:1.22

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /helloer

EXPOSE 8080

# install migrations
RUN go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Run
CMD ["/helloer"]