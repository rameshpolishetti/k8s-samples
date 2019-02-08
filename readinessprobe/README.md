# readinessProbe example

```
cd readinessprobe
minikube start
eval $(minikube docker-env)
docker build -t readiness:latest .

kubectl apply -f deployment.yml
kubectl apply -f service.yml

kubectl get pods
```
Access server application
```
minikube ip
curl $(minikube ip):30061
```

Loginto one pod and start server app in the port 8080
```
kubectl exec -it readiness-sample-XXXX /bin/bash
# go run server.go -port=8080
```

Access server application
```
minikube ip
curl $(minikube ip):30061
```
