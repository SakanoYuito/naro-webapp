FROM golang:latest

WORKDIR /
COPY ./ /

RUN go mod init github.com/SakanoYuito/naro-webapp \
  && go get github.com/labstack/echo/v4 \
  && go get gorm.io/gorm \
  && go get gorm.io/driver/mysql \
  && go mod tidy \
  && go build

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=arm64

EXPOSE 80

CMD ["go", "run", "main.go"]