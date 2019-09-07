spin-up:
	env GOOS=linux go build -o bin/wtf1
	docker build -t wtf1 .
	yes | docker system prune
	kubectl apply -f k8s/minikube/deployment.yml

spin-down:
	kubectl delete deploy wtf1