### Deploy httpserver

```sh
kubectl create ns myistio
kubectl label ns myistio istio-injection=enabled
kubectl create -f httpserver.yaml -n myistio
```

```sh
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=example Inc./CN=*.example.com' -keyout example.com.key -out example.com.crt
kubectl create -n istio-system secret tls example-credential --key=example.com.key --cert=example.com.crt
kubectl apply -f istio.yaml -n myistio
```

### Check ingress ip

```sh
kubectl port-forward service/istio-ingressgateway -n istio-system 443:443
```

### Access the httpserver via ingress

```sh

curl --resolve test.example.com:443:127.0.0.1 https://test.example.com/healthz -v -k
```
