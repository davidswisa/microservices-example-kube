# microservices-example-kube
microservices-example-kube



steps:
install kompose
$ go get -u github.com/kubernetes/kompose

follow:
https://kubernetes.io/docs/tasks/configure-pod-container/translate-compose-kubernetes/

install k0s
curl -sSLf https://get.k0s.sh | sudo sh





https://www.youtube.com/watch?v=aHPGf6FsY7Y
https://computingforgeeks.com/deploy-kubernetes-cluster-on-linux-with-k0s/





TODO:

hardcoded ip need to fix in ui/index.html and prod/main.go
http://10.144.74.86/


replace network api version:
apiVersion: networking.k8s.io/v1

replace all services image value:
- image: microservices-example-kube_ui:latest
imagePullPolicy: "IfNotPresent"