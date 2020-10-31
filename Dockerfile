FROM golang:1.14

WORKDIR /go/src/github.com/ssoyyoung.p/seoulbitz-Backend

ENV key=value

RUN go get -u github.com/go-sql-driver/mysql
RUN go get github.com/labstack/echo
RUN go get github.com/dgrijalva/jwt-go

CMD go run main.go
