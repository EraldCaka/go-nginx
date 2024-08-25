# go with nginx example

A small example of a basic application with single http request for nginx testing purposes

# Utilities

- creating services and using nginx as a load balancer utilizing round robin algorithm for horizontal scaling
- using pub/sub structure of watermill library. I used gochannel methods of pub/sub which is an in-memory solution but soon will use apache kafka or rabbitMQ for service-to-service communication