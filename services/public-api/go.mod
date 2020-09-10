module sg/services/public-api

go 1.13

require (
	github.com/Masterminds/squirrel v1.4.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-gorp/gorp v2.2.0+incompatible // indirect
	github.com/google/uuid v1.1.2
	github.com/gorilla/schema v1.1.0
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/lib/pq v1.7.0
	github.com/poy/onpar v1.0.0 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/sys v0.0.0-20200905004654-be1d3432aa8f // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v0.0.0-20200902210233-8630cac324bf // indirect
	gopkg.in/gorp.v2 v2.2.0

	libraries v0.0.0
)

replace libraries v0.0.0 => ../../libraries/golang
