# Exercise 3: Rolling deployments


## Now rolling deploy the new version of the service

Inspect the new `yaml`-file `deploy-rolling`, then execute:

    cd deploy-rolling
    kubectl apply -f deals-deploy-rolling.yaml  --namespace sock-shop

Change to the Kubernetes UI and switch in the *Namespaces* to *sock-shop*.
Then click on *Replica Sets* and select `deals`. You'll see how the pods are being replaced one-by-one.

## Make the frontend to display a random deal

Hook up the frontend to the deals backend via *JQuery* and choose a random one.
You can either solve this client-side via javascript or build a new version of the deals service that returns random entries.
(hint: there's a `front-end.patch` file that can act as a source of inspiration)

    docker build -t 12-factor-workshop/front-end:v2 .

Change the deployment from the frontend to use the new image. The deployment is here:

    cd ~/microservices-demo/microservices-demo/deploy/kubernetes/manifests
    micro front-end-dep.yaml

Now change the `image` and also tell Kubernetes to not try to pull it from Docker Hub:

        image: 12-factor-workshop/front-end:v2
        imagePullPolicy: Never

Finally, tell Kubernetes to roll out the new version:

    kubectl apply -f front-end-dep.yaml