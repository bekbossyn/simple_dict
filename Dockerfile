FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN sudo apt-get install git
RUN go get -u github.com/kataras/iris
RUN go get -u github.com/labstack/gommon/log
RUN go get -u github.com/lib/pq
RUN go build -o main .
CMD ["/app/main"]
