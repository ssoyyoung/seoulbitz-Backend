FROM golang:1.14

WORKDIR /go/src/github.com/ssoyyoung.p/seoulbitz-Backend

ENV key=value
ENV LC_ALL=C.UTF-8

RUN go get -u github.com/go-sql-driver/mysql
RUN go get github.com/labstack/echo
RUN go get github.com/dgrijalva/jwt-go

CMD go run main.go
