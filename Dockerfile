FROM golang:1.18-alpine

RUN mkdir /app

ENV DOMAIN_API_KEY=${DOMAIN_API_KEY}

ADD . /app

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]