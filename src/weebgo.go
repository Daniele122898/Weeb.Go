package weebgo

import (
	"github.com/Daniele122898/weeb.go/src/net"
	"github.com/Daniele122898/weeb.go/src/data"
)

func Authenticate(token string) error{
	return net.Authenticate(token)
}

func GetTags(hidden bool) (*data.TagsData, error) {
	return net.GetTags(hidden)
}

func GetTypes(hidden bool) (*data.TypesData, error) {
	return net.GetTypes(hidden)
}

func GetRandomImage(typ string, tags []string,filetype net.FileType,nsfw net.Nsfw, hidden bool) (*data.RandomData, error){
	return net.GetRandom(typ, tags, filetype, nsfw, hidden)
}
