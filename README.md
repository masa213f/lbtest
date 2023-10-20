# lbtest


## Apply server
```
kubectl apply -f lbtest-server.yaml
``````

## Apply Client and watch logs
```
kubectl apply -f lbtest-client.yaml
kubectl logs -n sandbox -l app=lbtest -f
kubectl delete -f lbtest-client.yaml
```

## Change the backend
```
$ date "+%Y-%m-%dT%H:%M:%S.%3N"; kubectl label -n sandbox pod -l app=lbtest backend-; date "+%Y-%m-%dT%H:%M:%S.%3N"
$ date "+%Y-%m-%dT%H:%M:%S.%3N"; kubectl label -n sandbox pod -l app=lbtest backend=yes; date "+%Y-%m-%dT%H:%M:%S.%3N"
``````
