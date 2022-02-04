FROM golang:1.17-alpine

WORKDIR /hackmanapi

ENV GO111MODULE=auto
ENV GIN_MODE=release

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /usr/local/hackmanapi .

CMD [ "/usr/local/hackmanapi" ]
