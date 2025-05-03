# mock/Dockerfile
FROM golang:1.23-alpine

WORKDIR /app
COPY . .

RUN go mod init publisher && go get github.com/eclipse/paho.mqtt.golang && go build -o publisher .

CMD ["./publisher"]
