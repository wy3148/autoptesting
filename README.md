# autoptesting setup guide

This is done by using a simple web framework-beego, also some other redis opensource package is used.

## prerequisite
Go (> 1.11)
redis single server (no need to be cluster)


```
export GOPATH=xxx
export autopilotapikey=your_api_key

```


```
go get github.com/wy3148/autoptesting
go run main.go
```

service will listening on localhost:8080

# config
github.com/wy3148/autoptesting/src/conf/app.conf


# api testing
contact/id GET
```
http://localhost:8080/contact/person_BF8A2F05-0E71-496F-A76E-AE99086CC6BC
```

contact POST
```
{
  "contact": {
    "FirstName": "Slarty122222",
    "LastName": "Bartfast233333",
    "Email": "test113334443@slarty.com"
  }
}
```

# testing
```
store_test.go unit test
tests/default_test.go (local endpoint testing)
more testing is to be added
```


