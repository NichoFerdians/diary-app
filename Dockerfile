FROM golang:alpine as builder

RUN apk update && apk add --no-cache git build-base

RUN mkdir /app

WORKDIR /app

# Caching go module
COPY go.mod .
COPY go.sum .

RUN go mod download
ADD . .

RUN go get github.com/githubnemo/CompileDaemon

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

EXPOSE 8080

CMD /wait \
  && CompileDaemon --build="go build main.go" --command=./main