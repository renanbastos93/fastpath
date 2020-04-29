# fastpath
based on urlpath and path from go


cases


For each segment
if optional && pathSegment == "" { break }
if param is wildcard or optional and is not last it must have value
else loop stops and returns

```
/api/:param/:opt?
/api/*
/api/:param/*
/api/:opt?
```

I created some use cases to routing, can you validate to me?

# Use Cases

## pattern: /api/v1/:param/*
``` bash
/api/v1/entity # sucess
/api/v1/entity/id # sucess
/api/v # error
/api/v2 # error
/api/v2 # error
```
## pattern: /api/v1/:param?
``` bash
/api/v1/ # sucess
/api/v1 # sucess
/api/v1/entity # sucess
/api/v # error
/api/v2 # error
/api/v2 # error
```
## pattern: /api/v1/*
``` bash
/api/v1 # sucess
/api/v1/ # sucess
/api/v1/entity # sucess
/api/v # error
/api/v2 # error
/api/v2 # error
```
## pattern: /api/v1/:param
``` bash
/api/v1/entity # sucess
/api/v1/ # error
/api/v1 # error
```
## pattern: /api/v1/const
``` bash
/api/v1/const # sucess
/api/v1 # error
/api/v1/ # error
/api/v1/something # error
```

## pattern: /api/v1/:param?/*
``` bash
# panic
```