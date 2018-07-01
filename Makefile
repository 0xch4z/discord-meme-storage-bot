BINARY := ./discord-meme-storage-bot

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix ncgo -o $(BINARY) .
