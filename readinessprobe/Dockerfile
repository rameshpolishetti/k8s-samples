FROM golang:1.11

RUN apt-get -y update &&\ 
    apt-get -y install lsof git

ADD server.go .

EXPOSE 8086 8080

# start the sample server application on port 8086
CMD go run server.go -port=8086