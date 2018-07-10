Basic REST API in Go

# Running locally (needs posgres running)
```
# set env vars
go run main.go
```

# Create a new pament record
```
curl -H "Content-Type: application/json" -X POST -d @payment.json http://127.0.0.1:3000/payments
```

# Get version
```
curl http://127.0.0.1:3000/version
```

# Get payments
```
curl http://127.0.0.1:3000/payments
```

# Build
```
GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/achanda/testrest/version.Version=`git rev-parse HEAD`" -o testrest .
```

# Generate image
```
docker build . -t testrest
```
