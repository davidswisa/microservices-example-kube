# microservices-example-kube

## About the project:  

This simple reservation application uses a queue manager for all reservation tasks like create, update, and delete.
And use a more direct approach to get the data from the DB.

### How it all works

```mermaid
graph TD;
GUI(GUI) -->|GET| Querier;
GUI(GUI) -->|POST,PUT,DELETE| Producer;

Querier(Querier) --> ORM;
Kafka --> Producer;
Kafka --> Consumer;
ORM --> Database;
Consumer --> ORM;

```

### Service Table
| Container | Port | Purpose  |
| :---      | :-   | :- |
| ui        | 8084 | gui |
| cons      | --   | get messages (tasks) from kafka and forwards to orm |
| orm       | 5431 | used as a mitigator for DB |
| prod      | 8080 | receive http request from ui and create messages (tasks) via kafka |
| querier   | 8081 | receive http request from ui and forwards a query via orm |
| kafka     | 9092, 8082, 8083 | message bus |
| zookeeper | 2181 | ? |
| db        | 5432 | data base |


## Prerequisite:

### Docker-compose 

In order to simplify the containers build process we decided to use `docker-compose` make sure to have it installed.  

### kubernetes:

For the services we used `k3s` with docker as the container runtime

``` bash
curl -sfL https://get.k3s.io | sh -s - --docker
```

configure local kubctl to work with container k3s  

``` bash  
echo "export KUBECONFIG=/etc/rancher/k3s/k3s.yaml" >> ~/.bashrc
```

## Getting Started:

```bash
make run
```
