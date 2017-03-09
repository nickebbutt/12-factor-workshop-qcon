# Exercise 2: Building, Deploying, Extending


## Build a docker image

    cd ~/12-factor-workshop
    cd deploy-deals-service

Inspect the `Dockerfile` here, this builds your service.
Create (or copy) a similar `Dockerfile` in your code directory and build the image.
If you need inspiration for your own language, browse the https://github.com/microservices-demo organisation for 
services written in your language of choice, they all come with a `Dockerfile`.

    cp ../simple-deals-service/main.go .
    docker build -t 12-factor-workshop/deals:v1 .

You can see the new images with `docker images ls`

    docker images | grep deals

Try it out:

    # run the container as name "deals" as daemon, expose port 8888 and remove it again when it stops
    docker run -d --rm -p 8888:8888 --name deals 12-factor-workshop/deals:v1

    # check if it works
    curl localhost:8888/deals?id=1

    # stop the container
    docker stop deals

## Deploy the Docker image to Kubernetes

Inspect the `yaml` files in this directory. They are used to deploy the image.

First, we create a Kubernetes *Deployment* that rolls out the *Pods* with the containers and keep them up by creating a *ReplicaSet*

    kubectl create --namespace sock-shop -f deals-dep.yaml

Inspect the Kubernetes UI if it was created.
Then create a service to make the deals available:

    kubectl create --namespace sock-shop -f deals-svc.yaml

You can either check the UI if the service was created successfully, or employ `kubectl`

    kubectl describe svc deals --namespace sock-shop

You will see a line `IP: 10.xx.xx.xx`. Try to `curl` the service with it:

    curl 10.[ip.from.kubectl]:8899/deals?id=1

## OpenAPI / Swagger

The existing services all come with their API specification, e.g.

    https://github.com/microservices-demo/orders/blob/master/api-spec/orders.json

Or 
    https://github.com/microservices-demo/user/blob/master/apispec/user.json

For our new *Deals* service, the api will be quite simple. Extend the API spec to include a `description` field.

    cd add-description-field
    # edit deals-swagger.yaml

Then, generate a sample implementation with `swagger-codegen` and your favourite language (`swagger-codegen` is already installed on the VM)

    swagger-codegen generate -i ./deals-swagger.yaml -l nodejs-server -o deals-nodejs

Now inspect the files generated in `deals-nodejs`

Bonus-material, if you have time: Check out how you can use `dredd` to automate tests against the API specification, by inspecting the contents of `~/microservices-demo/microservices-demo/openapi/`.

## Add a data store to the service

We'll store deal in a MongoDB. For this, we need to create one first, see the `adding-a-datastore` directory.

There are two additional manifests that need to be sent to Kubernetes (`deals-db-svc.yaml` and `deals-db-dep.yaml`).

    docker build -t 12-factor-workshop/deals:v3 .
