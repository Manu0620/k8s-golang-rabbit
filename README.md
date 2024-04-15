# Deploy apps to docker hub

docker login

docker build . -t tuusuariodedocker/nombredelaapp:latest

docker push tuusuariodedocker/nombredelaapp:latest

# Get started

minikube start

kubectl config use-context minikube

kubectl create ns rabbits

# Deploy

kubectl apply -n rabbits -f .\kubernetes\rabbit-rbac.yaml

kubectl apply -n rabbits -f .\kubernetes\rabbit-configmap.yaml

kubectl apply -n rabbits -f .\kubernetes\rabbit-secret.yaml

kubectl apply -n rabbits -f .\kubernetes\rabbit-statefulset.yaml

# Acces to rabbit

kubectl -n rabbits port-forward rabbitmq-0 8080:15672

Go to htttp://localhost:8080

Username: guest

Password: guest

# Deploy apps

cd nombredelacarpetadelaapp

kubectl apply -n rabbits -f nombredelarchivo.yaml

# Deploy postgres to k8s

kubectl apply -n rabbits -f postgres-configmap.yaml

kubectl -n rabbits get configmap

kubectl apply -n rabbits -f psql-pv.yaml

kubectl apply -n rabbits -f psql-claim.yaml

kubectl -n rabbits get pv

kubectl -n rabbits get pvc

kubectl apply -n rabbits -f ps-deployment.yaml

kubectl -n rabbits get deployments

kubectl -n rabbits get pods

kubectl -n rabbits apply -f ps-service.yaml

kubectl -n rabbits get svc

kubectl exec -n rabbits -it podname -- psql -h localhost -U username --password -p 5432 dbname# Deploy apps to docker hub

docker login

docker build . -t tuusuariodedocker/nombredelaapp:latest

docker push tuusuariodedocker/nombredelaapp:latest

# Get started

minikube start

kubectl config use-context minikube

kubectl create ns rabbits

# Deploy

kubectl apply -n rabbits -f .\kubernetes\rabbit-rbac.yaml

kubectl apply -n rabbits -f .\kubernetes\rabbit-configmap.yaml

kubectl apply -n rabbits -f .\kubernetes\rabbit-secret.yaml

kubectl apply -n rabbits -f .\kubernetes\rabbit-statefulset.yaml

# Acces to rabbit

kubectl -n rabbits port-forward rabbitmq-0 8080:15672

Go to htttp://localhost:8080

Username: guest

Password: guest

# Deploy apps

cd nombredelacarpetadelaapp

kubectl apply -n rabbits -f nombredelarchivo.yaml

# Deploy postgres to k8s

kubectl apply -n rabbits -f postgres-configmap.yaml

kubectl -n rabbits get configmap

kubectl apply -n rabbits -f psql-pv.yaml

kubectl apply -n rabbits -f psql-claim.yaml

kubectl -n rabbits get pv

kubectl -n rabbits get pvc

kubectl apply -n rabbits -f ps-deployment.yaml

kubectl -n rabbits get deployments

kubectl -n rabbits get pods

kubectl -n rabbits apply -f ps-service.yaml

kubectl -n rabbits get svc

kubectl exec -n rabbits -it podname -- psql -h localhost -U username --password -p 5432 dbname
