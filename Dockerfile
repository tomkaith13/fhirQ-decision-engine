FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /main src/main.go
RUN ls -la

EXPOSE 8080

CMD [ "/main" ]