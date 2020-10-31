FROM golang:1.14

WORKDIR /go/src/github.com/ssoyyoung.p/seoulbitz-Backend

ENV key=value

RUN go get -u github.com/go-sql-driver/mysql
RUN go get github.com/labstack/echo

CMD go run main.go