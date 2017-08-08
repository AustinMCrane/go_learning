FROM golang
ADD . /go/src/github.com/AustinMCrane/wedding_api
WORKDIR /go/src/github.com/AustinMCrane/wedding_api
RUN ls
RUN go build
ENTRYPOINT ./wedding_api
EXPOSE 8888