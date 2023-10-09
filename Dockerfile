FROM golang:1.20

WORKDIR /app

COPY main.go .

RUN go build main.go

CMD [ "./main" ]