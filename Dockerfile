FROM golang:1.25.0

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
COPY main.go .
COPY Controllers/ Controllers/
COPY Models/ Models/
COPY Routes/ Routes/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ping_app
ENTRYPOINT ["/app/ping_app"]