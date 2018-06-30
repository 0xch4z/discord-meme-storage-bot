FROM golang:latest
LABEL maintainer="Charles Kenney <me@chaz.codes>"

ENV GOBIN /go/bin

WORKDIR /go/src/github.com/Charliekenney23/discord-meme-bot

COPY . .

RUN mkdir /usr/share/memes
VOLUME "/usr/share/memes"

RUN dep ensure

RUN make build
CMD ["./discord-umad-bot"]

EXPOSE 3000
