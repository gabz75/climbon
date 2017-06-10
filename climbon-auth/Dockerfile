FROM golang

ADD . /go/src/github.com/gabz75/climbon-auth

RUN go get github.com/gin-gonic/gin
RUN go get gopkg.in/appleboy/gin-jwt.v2

RUN go install github.com/gabz75/climbon-auth

ENTRYPOINT /go/bin/climbon-auth

EXPOSE 8080

CMD ["go", "run", "climbon-auth"]
