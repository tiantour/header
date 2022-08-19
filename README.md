# header

get data from grpc-gateway header

set set to grpc-gateway header

# how to use

1. get

```golang
key := "foo"
value := header.NewRequest().Get(ctx, key)
```

2. set

```golang
data := []string{"food", "bar"}
ctx = header.NewResponse().Set(ctx, data...)
```