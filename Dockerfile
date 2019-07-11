FROM golang:1.12.2 as builder
COPY . /go/github.com/tmarcus87/slackutil
WORKDIR /go/github.com/tmarcus87/slackutil
ENV GO111MODULE=on
RUN go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' -o slackutil main.go

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/github.com/tmarcus87/slackutil/slackutil /usr/local/bin/slackutil
ENTRYPOINT [ "slackutil" ]
