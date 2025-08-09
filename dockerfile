FROM golang:1.23

WORKDIR /app

COPY  go.mod .
COPY  go.sum .
COPY  main.go .

COPY ./ .


RUN go mod download

ENV Logger_Channel="Logger"

RUN go build -o bin .
#chmod +x /path/to/executable

#ENTRYPOINT [ "/app/LoggerInject.exe" ]
ENTRYPOINT [ "/app/bin" ]
