FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o location-project .

EXPOSE 8081

CMD [ "./location-project" ]