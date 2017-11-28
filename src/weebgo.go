package weebgo

import "github.com/Daniele122898/weeb.go/src/net"

func Authenticate(token string) error{
	return net.Authenticate(token)
}
