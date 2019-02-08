# readinessProbe example

This example demonstrates how to use redinessProbe feature.

## Steps

1. Build docker image and deploy minikube based k8s cluster
```
cd readinessprobe
minikube start
eval $(minikube docker-env)
docker build -t readiness:latest .
kubectl apply -f deployment.yml
kubectl apply -f service.yml
kubectl get pods
```

2. Try to access server application running inside a pod 
```
minikube ip
curl $(minikube ip):30061
```
you should see error message like 
```
curl: (7) Failed to connect to 192.168.99.100 port 30061: Connection refused
```

3. Loginto any one of the pod and start sample server application on the port 8080, to indicate pods readiness
```
kubectl exec -it readiness-sample-XXXX /bin/bash
# go run server.go -port=8080
```

4. Access server application again
```
minikube ip
curl $(minikube ip):30061
```
You should see response from the server like
```
{'message':'hello from server', 'timestamp':'2019-02-08 19:22:47.642158446 +0000 UTC m=+698.008702716'}
```
