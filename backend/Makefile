build:
	GOOS=linux go build -o app
	docker build -t gcr.io/hazel-champion-200108/say .
	rm -f app
run:
	docker run --rm -v $(shell pwd)/data:/data -p 8080:8080 -w /data maxim/say
play:
	afplay data/output.wav
list:
	gcloud container clusters list
push:
	gcloud docker -- push gcr.io/hazel-champion-200108/say
cred:
	gcloud container clusters get-credentials cluster-maxim
nodes:
	kubectl get nodes
create-deployment:
	kubectl create -f kubernetes.yaml
svc:
	kubectl get svc
watch:
	kubectl get svc -w
