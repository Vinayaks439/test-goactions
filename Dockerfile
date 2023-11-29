FROM golang:1.14.2-alpine3.11
RUN apk add git

COPY . /home/src
WORKDIR /home/src
RUN go build -o /bin/action ./

ENTRYPOINT [ "/bin/action" ]