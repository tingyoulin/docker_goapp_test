FROM golang

WORKDIR /

COPY . .

RUN go get github.com/go-redis/redis
RUN go get google.golang.org/grpc

RUN go build -o main ./

CMD ["./main"]