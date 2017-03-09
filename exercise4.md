# Exercise 4: Metrics, Logging, Tracing


## Add log collection with fluentd and Kibana

Let's first look at how can ensure in the code, that logging works properly

    -> Ian

Try to add this to your own service, look at the `logging-and-metrics` directory.

If you have added some code, build a new image:

    docker build -t 12-factor-workshop/deals:v4 .

Now, we need to add a few more moving parts to collect all logs with *fluentd*, 
index them in *elasticsearch* and make them searchable with *Kibana*:

    cd ~/microservices-demo/microservices-demo/deploy/kubernetes/manifests
    kubectl create -f kubernetes/manifests-logging/elasticsearch.yml
    kubectl create -f kubernetes/manifests-logging/fluentd-daemon.yml
    kubectl create -f kubernetes/manifests-logging/kibana.yml

You can now access the Kibana dashboard via

    http://[ip-of-your-vm]:31601

After creating the initial index, you can click around in the socks shop, and 
soon you will see logs appearing in the discover view of Kibana.

## Add additional metrics

Useful metrics can often be exposed directly by the web framework.
Also Prometheus brings a very useful library to do the heavy lifting with maintaining counters etc.
For an example in Java, see:

    https://github.com/microservices-demo/carts/blob/master/src/main/java/works/weave/socks/cart/monitoring/PrometheusMetricWriter.java

In go, with `go-kit` this, is achieved with "middlewares". Both metrics and logging:

    https://github.com/microservices-demo/payment/blob/master/middlewares.go#L49

Prometheus even provides it's own handler for the scraping endpoint:

    https://github.com/microservices-demo/user/blob/master/api/transport.go#L109

For more inspiration, see also http://container-solutions.com/microservice-monitoring-with-prometheus/

To be able to collect metrics, we can install Prometheus:

    cd logging-and-metrics
    kubectl apply -f prometheus_graphana.yaml

If you reload the Kubernetes dashboard, you will see a new *monitoring* namespace.

You can access the UI via:

    http://[ip-of-your-vm]:31300/login
    # admin / admin

As alternative, there's a brand-new "Operator" available for Prometheus from CoreOS.

## Add tracing

With many web frameworks, Tracing via Zipkin can be added relatively easy be integrating it with the request/response cycle:

    https://github.com/microservices-demo/user/blob/master/main.go#L57

Zipkin is already installed along with the Sock Shop. 
You can reach the UI by opening

    http://[ip-of-your-vm]:30002

Tracing isn't yet added properly to all the services, but should be able to see e.g. the ones includin
the user service here:

    http://[ip-of-your-vm]:30002/?serviceName=user&spanName=all

Without any code changes, tools like Weave Scope can show you how individual services interact:

    cd logging-and-metrics
    kubectl apply -f weavescope.yaml

Browse to:

    http://[ip-of-your-vm]:30040/

You can change the namespace in the lower left corner of the screen.