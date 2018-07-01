FROM golang:1.10 AS builder
LABEL maintainer="Charles Kenney <me@chaz.codes>"

ENV GOPATH /go
ENV GOBIN $GOPATH/bin
RUN PATH=$GOBIN:$PATH

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR /go/src/github.com/Charliekenney23/discord-meme-storage-bot

COPY . .

RUN mkdir /usr/share/memes /usr/share/db
RUN touch /usr/share/db/dmsb.db
VOLUME ["/usr/share/memes", "/usr/share/db"]

RUN dep ensure --vendor-only

RUN make build
CMD ./discord-meme-storage-bot
