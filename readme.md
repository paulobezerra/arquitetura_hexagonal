#Install Mockgem

`go install github.com/golang/mock/mockgen@v1.6.0`

`export PATH=$PATH:$(go env GOPATH)/bin`

`mockgen -source application/product.go -destination application/mocks/application.go`

#Install Cobra

`go install github.com/spf13/cobra-cli@latest`