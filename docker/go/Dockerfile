FROM golang:1.13.4-buster

WORKDIR /go/src/github.com/blacknikka/go-k8s

RUN go get -u github.com/jinzhu/gorm \
    && go get -u github.com/go-sql-driver/mysql

ENV PATH $PATH:/go/bin

CMD ["go", "run", "./main.go"]
