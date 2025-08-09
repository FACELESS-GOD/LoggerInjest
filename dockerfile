FROM golang:1.23

WORKDIR /app

COPY  go.mod .
COPY  main.go .

ENV Logger_Channel="Logger"

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]