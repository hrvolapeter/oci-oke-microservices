FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o /service-b

CMD [ "/service-b" ]