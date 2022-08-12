FROM golang:1.19-alpine

ENV ROOT=/go/src/Linebot

RUN mkdir ${ROOT}

WORKDIR ${ROOT}

COPY go.mod ./

RUN go mod download

RUN apk upgrade --update && \
    apk --no-cache add git

#EXPOSE 19000

# CMD ["realize", "start", "--build","--run"]
