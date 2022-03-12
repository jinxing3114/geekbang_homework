# Build docker image

```sh
docker build -t mywebser:v3.1 .

```

## deployment app

```sh
kubectl apply -f myweb.yaml
kubectl get pods myweb -o yaml
```