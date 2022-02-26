# ConfigMap

```sh
#configmap from file
kubectl create -f configMap.yaml
kubectl get configmap -oyaml myweb-config
```

## pod

```sh
kubectl create -f myweb.yaml
kubectl get pods myweb -o yaml
```