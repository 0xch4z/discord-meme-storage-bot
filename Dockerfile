FROM golang:latest
LABEL maintainer="Charles Kenney <me@chaz.codes>"

ENV GOBIN /go/bin

WORKDIR /go/src/github.com/Charliekenney23/discord-meme-storage-bot

COPY . .

RUN mkdir /usr/share/memes /usr/share/db
RUN /usr/share/db/dmsb.db
VOLUME ["/usr/share/memes", "/usr/share/db"]

RUN dep ensure

RUN make build
CMD ["./discord-meme-storage-bot"]

EXPOSE 3000
