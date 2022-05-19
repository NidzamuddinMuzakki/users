FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download
COPY . ./

RUN go build -o /docker-gs-ping

EXPOSE 80 443

CMD [ "/docker-gs-ping" ]