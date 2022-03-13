# Build docker image

```sh
docker build -t mywebser:v3.1 .

```

## deployment app

```sh
kubectl apply -f myweb.yaml
kubectl get pods myweb -o yaml
```

## monitoring data
![image](https://raw.githubusercontent.com/jinxing3114/geekbang_homework/master/cloud_native/week_10/panel.png)