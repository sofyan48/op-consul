## SERVICE DISCOVERY EXAMPLE
Service Discovery Example By Consul

## HOW TO
using go mod and golang 1.13 or latest
Install Consul
**Mac Os**
Installing from brew
```
brew install consul
```
**Linux**
Download Consul binary
```
wget https://releases.hashicorp.com/consul/1.6.2/consul_1.6.2_linux_amd64.zip
```
Extract zip
```
unzip consul_1.6.2_linux_amd64.zip
```
Copy to /usr/local/bin
```
cp consul /usr/local/bin
```

### Run Consul
Run Consul 
```
consul agent -dev -ui
```

### Run Service Product
```
go run product/main.go
```

### Run Service Order
```
go run order/main.go
```

### Run Service User
```
go run user/main.go
```

## Example
Get Product from user
```
curl -XGET localhost:8080/user/product
```

Get Orders from user
```
curl -XGET localhost:8080/user/orders
```