FROM golang:1.8

WORKDIR /go/src/github.com/austinmcrane/wedding_app_api
COPY . .

ENV PORT 8888
EXPOSE 8888
EXPOSE 3000

RUN go get github.com/codegangsta/gin
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD gin run
