.PHONY: run-compose run build clean

run-compose:
	docker-compose up -d --build

run:
	kubectl apply -f ./kube

build:
	docker rmi $$(docker images | grep 'microservices-example-kube') || true
	docker-compose build --no-cache

.PHONY: rebuild-%
rebuild-%:
	docker rmi "microservices-example-kube-$*"
	docker-compose build --no-cache "$*"

clean: kube-clean
	docker rmi $$(docker images | grep 'microservices-example-kube') || true

kube-clean:
	kubectl delete deployments --all
	kubectl delete services --all
	kubectl delete pods --all
	kubectl delete daemonset --all
	kubectl delete NetworkPolicy --all
	kubectl delete replicaset --all
