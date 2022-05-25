# microservices-example-kube
microservices-example-kube



steps:
We used Kompose in order to convert the majority of the docker-compose file conveniently.

Installing Kompose:
```bash
go get -u github.com/kubernetes/kompose
```

follow:
https://kubernetes.io/docs/tasks/configure-pod-container/translate-compose-kubernetes/

install k3s with docker as the container runtime:
```bash
curl -sfL https://get.k3s.io | sh -s - --docker
```

### Changes Notes:

- Replace network api version to the default one:
```yaml
apiVersion: networking.k8s.io/v1
```

- Replace all services image value:
    - After using `docker-compose build` this change allows pulling from the local docker registry (when using docker as the container runtime)
```yaml
- image: microservices-example-kube_ui:latest
  imagePullPolicy: "IfNotPresent"
```

- Add to deployment the port to expose out to the host:
    - This allows access to the pods from the host machine.
```yaml
hostPort: 8080
```
