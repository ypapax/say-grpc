build:
	GOOS=linux go build -o app
	docker build -t maxim/say .
	rm -f app
run:
	docker run --rm -v $(shell pwd)/data:/data -p 8080:8080 -w /data maxim/say
play:
	afplay data/output.wav