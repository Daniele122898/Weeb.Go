package weebgo

import (
	"github.com/Daniele122898/weeb.go/src/net"
	"github.com/Daniele122898/weeb.go/src/data"
)

func Authenticate(token string, typ data.TokenType) error{
	return net.Authenticate(token, typ)
}

func GetTags(hidden bool) (*data.TagsData, error) {
	return net.GetTags(hidden)
}

func GetTypes(hidden bool) (*data.TypesData, error) {
	return net.GetTypes(hidden)
}

func GetRandomImage(typ string, tags []string,filetype data.FileType,nsfw data.Nsfw, hidden bool) (*data.RandomData, error){
	return net.GetRandom(typ, tags, filetype, nsfw, hidden)
}
