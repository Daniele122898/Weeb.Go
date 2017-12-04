package data

type WelcomeData struct{
	Version string
	Status int
	Message string
}

type TypesData struct{
	status int
	Types []string
}

type TagsData struct{
	status int
	Tags []string
}

type Tags struct{
	Name string
	Hidden bool
	User string
}

type RandomData struct{
	Id string
	BaseType string
	FileType string
	MimeType string
	Account string
	Hidden bool
	Nsfw bool
	Tags []Tags
	Url string
}

type FileType int
type Nsfw int
type TokenType int

const(
	BEARER TokenType = iota
	WOLKE
)

const(
	JPG FileType = iota
	PNG
	GIF
	ANY
)

const(
	FALSE Nsfw = iota
	TRUE
	ONLY
)


