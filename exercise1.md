# Exercise 1: The Socks Shop and Deals


## Explore the Sock Shop code

    cd microservices-demo/microservices-demo
    # this was cloned from https://github.com/microservices-demo/microservices-demo

## Run the socks shop

Deploy to the local Kubernetes cluster

    cd simple-deals-service
    kubectl create -f complete-demo.yaml

Check if pods ("containers") are running

    kubectl get pods --namespace sock-shop

Check if services that the pods accessible are up

    kubectl get services --namespace sock-shop

## Inspect what Kubernetes has started

Browse to

    http://[ip-of-your-vm]:8001/ui/

Change to namespace `socks-shop` in the drop down (top left)

If you can't access this URL, you might need to execute `kubectl proxy --address=0.0.0.0 --port=8001  --disable-filter=true &`

## Browse the Socks shop

Point your browser to: http://[ip-of-your-vm]:30001/

You can also login using `user` and `password`.

## Get familiar with a few of the services

There are services in Java, C#, go and nodejs. 
Like all good 12-Factor-Citizens, they include everything needed to be built, e.g. a Travis-CI config:

    https://github.com/microservices-demo/orders/blob/master/.travis.yml

## Check out the frontend

We want to display our deals later in the frontend, so let's check that out, too

    cd ~/microservices-demo
    git clone https://github.com/microservices-demo/front-end.git

Inspect the code, to see how services are currently integrated and where we should put the deals stuff later:

    front-end/public/topbar.html
    front-end/public/js/client.js
    front-end/public/js/front.js
    front-end/api/endpoints.js

## Write a very simple deals service

Requirements:

* Has at least 2 predefined deals:
  1. Buy 400 pairs, get one unmatched sock free!
  2. Free shipping anywhere in the Andromeda Galaxy
* Responds to a `HTTP` `GET` against `/deals?id=[id]`, e.g. `GET /deals?id=1`
* Adheres to basic 12 factor principles (logs to stdout, doesn't write files...)

Done? Compare with the example provided in Golang.

    cd simple-deals-service
    go run main.go &

    # check if it works
    curl localhost:8888/deals?id=1
    curl localhost:8888/deals?id=2

    # stop it again
    kill %1
