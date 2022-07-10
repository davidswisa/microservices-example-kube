.PHONY: run-compose run build clean

run-compose:
	docker-compose up -d --build
	# docker-compose up -d --build db kafka zookeeper

run: build
	kubectl apply -f ./kube

build:
	docker-compose build --no-cache

.PHONY: rebuild-%
rebuild-%:
	docker-compose build --no-cache "$*"

clean: kube-clean

kube-clean:
	kubectl delete deployments --all
	kubectl delete services --all
	kubectl delete pods --all
	kubectl delete daemonset --all
	kubectl delete NetworkPolicy --all
	kubectl delete replicaset --all
