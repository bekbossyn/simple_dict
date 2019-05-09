FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

#FROM golang:alpine
#RUN mkdir /app
#ADD . /app
#WORKDIR /app
#RUN apk add git
##RUN go get -u github.com/kataras/iris
##RUN go get -u github.com/labstack/gommon/log
##RUN go get -u github.com/lib/pq
#RUN go build -o main .
#CMD ["/app/main"]
