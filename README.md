# kubernetes-training

##### Deploy sample server in kubernetes cluster



- Build Docker image and push it to the Docker registry

```
docker build -t <imagename:tag> <path_to_Dockerfile>
docker push <docker-registry>/<imagename:tag>
```

- Add the image name to the deployment

```
containers:
        - image: <image-name>
```

- Create a namespace in kubernetes cluster

```
kubectl create -f nameSpace.yaml  or
kubectl create namespace <namespace-name>
```

- Create a Deployment

```
kubectl create -f deployment.yaml 
```

- Expose the Deployment to get external traffic into the cluster

    i. Expose with NodePort service
    
    - Create the service

    ```
    kubectl create -f serviceNodePort.yaml 
    ```
    - Get and describe the service
    ```
    kubectl get svc -n <namespace>
    kubectl describe svc <service name> -n <namespace>
    ``` 
    -Access the server with service
    ```
    curl http://<nodeport-ip>:port
    ``` 
    
    ii. Expose with LoadBalancer service
    
    - Create the service
    
     ```
     kubectl create -f serviceLoadBalancer.yaml  or
     kubectl expose deployment <deployment-name> --type=LoadBalancer --name=<service-name>
     ```
    - Access the server with service
      
     ```
     curl http://<external-ip>:port
     ```
      
- Check the cluster status

    i. Get and describe deployments

    ```
    kubectl get deployment -n <namespace>
    kubectl describe deployment <deployment-name> -n <namespace>
    ``` 
    
    ii. Get and describe pods
    
    ```
    kubectl get pods -n <namespace>
    kubectl describe pods -n <namespace>
    ``` 
    
    iii. Access the pods
    
    ```
    kubectl exec -it <pod name> -n <namespace> -- sh
    ``` 
- Clean up

    i. Delete Deployment
    
    ```
    kubectl delete deployment <deployment name> -n <namespace>
    ``` 
    
    ii. Delete Service
    
    ```
    kubectl delete svc <service name> -n <namespace>
    ``` 
    
    iii. Delete Namespace
    
    ```
    kubectl delete namespace <namespace>
    ``` 
    