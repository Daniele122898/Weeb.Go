package data

type WelcomeData struct{
	Version string
	Status string
	Message string
}

type TypesData struct{
	status string
	Types []string
}

type TagsData struct{
	status string
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


