FROM golang:latest

RUN go get -u github.com/gin-gonic/gin && \
    go get -u github.com/golang/dep/cmd/dep && \
    go get -u github.com/lib/pq && \
    go get -u github.com/jinzhu/gorm && \
    go get -u github.com/joho/godotenv && \
    go get -u github.com/go-redis/redis && \
    go get -u github.com/oxequa/realize

WORKDIR /go/src/github.com/go_gin_postgresql/