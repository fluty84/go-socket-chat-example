FROM golang
EXPOSE 8080
WORKDIR go/server
COPY ./ .
CMD go get github.com/gorilla/websocket && cd server && go run main.go