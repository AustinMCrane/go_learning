FROM golang:1.8

WORKDIR /go/src/github.com/austinmcrane/wedding_app_api
COPY . .

EXPOSE 8888
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"]
