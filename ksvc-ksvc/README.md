# Service to Service calling in knative

This guide has the steps to call a knative service from another.

## Step 1 - Set up the cluster local gateway

This is a one time process. Set up the istio cluster local gateway (if not set up already)

[Guide](https://knative.dev/docs/install/installing-istio/#updating-your-install-to-use-cluster-local-gateway)

## Step 2 - Call the service through cluster local gateway

If ```svc1``` is calling ```svc2``` in namespace ```abc``` Below is the url to call it.

```
http://svc2.abc.svc.cluster.local
```

## Step 3 (optional) - making backend svc cluster local

Incase if you want svc2 to be local to the cluster and block external access.

```
kubectl label ksvc svc2 serving.knative.dev/visibility=cluster-local
```

### Running this example

This example has 2 services 

1. test-cluster-local-svc
2. test-ingress-svc

ingress svc calls the local svc.

```
ko apply -f k8.yaml
```

