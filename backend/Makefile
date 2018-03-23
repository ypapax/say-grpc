build:
	GOOS=linux go build -o app
	docker build -t maxim/say .
	rm -f app
run:
	docker run --rm -v $(shell pwd)/data:/data -w /data maxim/say hello
play:
	afplay data/output.wav