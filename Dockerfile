FROM golang:alpine

RUN mkdir /app
COPY . /app/

WORKDIR /app

RUN go build -o cloudflare-dyndns

CMD ["/app/cloudflare-dyndns"]