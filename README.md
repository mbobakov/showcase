# simple swagger application

# init
- basic init
```
swagger init spec \
  --title "A clickhouse-swagger application" \
  --description "Shows how swagger and Clickhouse works together" \
  --version 1.0.0 \
  --scheme http \
  --consumes application/io.goswagger.metrics-app.v1+json \
  --produces application/io.goswagger.metrics-app.v1+json
```
- add operations and models to the spec
- run "swagger generate server" (goswagger version is 0.19)
- write service layer