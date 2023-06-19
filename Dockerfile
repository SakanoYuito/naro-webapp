FROM golang:latest

WORKDIR /app
COPY ./app /app

RUN go mod tidy \
  && go build

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=arm64

EXPOSE 80

CMD ["go", "run", "main.go"]