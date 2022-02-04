FROM golang:1.17-alpine

WORKDIR /hackmanapi

ENV GO111MODULE=auto

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /usr/local/hackmanapi .

CMD [ "/usr/local/hackmanapi" ]
