FROM golang:1.18

WORKDIR /go/apps/cicd-practice

COPY ./ ./

CMD ["go","run","main.go"]
