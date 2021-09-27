module github.com/devgek/mygrpc

go 1.15

replace mgithub.com/devgek/mygrpc => ../mygrpc

require (
	github.com/go-kit/kit v0.12.0
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)
