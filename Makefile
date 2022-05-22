run:
	docker-compose up -d --build
	# docker-compose up -d --build db kafka zookeeper  

kube-gen:
	rm -rf kube/*
	cd kube && kompose convert -f ../docker-compose.yml

run-kube:
	# kubectl apply $(ls test*.yaml | awk ' { print " -f " $1 } ')
	kubectl apply $$(ls test*.yaml | awk ' { print " -f " $$1 } ')

kube0-run:
	k0s kubectl apply -f kube

kube-clean:
	kubectl delete deployments --all
	kubectl delete services --all
	kubectl delete pods --all
	kubectl delete daemonset --all

kube0-clean:
	k0s kubectl delete deployments --all
	k0s kubectl delete services --all
	k0s kubectl delete pods --all
	k0s kubectl delete daemonset --all

run-prod:
	docker stop prod
	docker-compose up -d --build prod

run-cons:
	docker stop cons
	docker-compose up -d --build cons

run-orm:
	docker stop orm
	docker-compose up -d --build orm

run-ui:
	docker stop ui
	docker-compose up -d --build ui

run-querier:
	docker stop querier
	docker-compose up -d --build querier

prod:
	go run ./prod/main.go

cons:
	go run ./cons/main.go

querier:
	go run ./querier/main.go

orm:
	go run ./orm/main.go

test:
	scripts/send_request.sh
