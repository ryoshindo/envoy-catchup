## rewrite

Here is a sample using envoy's `path_rewrite`.

```shell
curl localhost:8081/app1/ready
app1 # response

curl localhost:8082/app2/ready
app2 # responseg
```

![](./doc/architecture.drawio.svg)
