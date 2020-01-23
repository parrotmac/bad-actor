FROM golang:latest

ENV HTTP_PORT 9090
EXPOSE 9090

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./
RUN go build -v ./

RUN mkdir /go/src/app/workdir

ENTRYPOINT ["/go/src/app/app"]
