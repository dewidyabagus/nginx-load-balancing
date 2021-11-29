FROM golang:alpine3.14

WORKDIR /app/webservice

COPY . .

EXPOSE 8000

RUN go build -o main

CMD ["./main"]