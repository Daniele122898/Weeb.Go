<hr/>
<h1 align="center">
	Weeb.Go
</h1>
<p align="center">
    A fast and easy to use weeb.sh API wrapper.
    <br>
	Originally build with Discord bots in mind like <a href='https://github.com/Daniele122898/GoBot'>GoBot</a>!
	<br>
	Join his Discord Server for support on the wrapper.
	<br>
	<br>
    <a href="https://discord.gg/Pah4yj5">
        <img src="https://discordapp.com/api/guilds/281589163659362305/widget.png?style=banner2">
    </a>
</p>
<hr/>

# Installation
## Go Get
Just use `go get https://github.com/Daniele122898/weeb.go`

### Dependencies
No extra packages needed.

### Token
To use the weeb.sh services you need an authentication token. You can request one at devs@weeb.sh

# Usage
`hidden` should be set as false by default. These only benefit uploaders and are useless otherwise!
### Authentication
First you have to authenticate once by passing your token. This has to be done before any further usage can occur and only has to be done once.
```go
//...
func Auth() error{
    err := weebgo.Authenticate("YOUR WEEB.SH TOKEN HERE")
    if err == nil {
        //DO ERROR HANDLING IN HERE
        return err
    }
    //No further logic needed...
    return err
}
//...
```

### Getting all available tags
```go
//...
func GetTags(hidden bool) (*data.TagsData, error){
    td, err := weebgo.GetTags(hidden)
    if err != nil{
        //error handling
        return nil, err
    }
    return td, nil
}
//...
```

### Getting all available types
```go
//...
func GetTypes(hidden bool) (*data.TypesData, error){
    td, err := weebgo.GetTypes(hidden)
    if err != nil{
        //error handling
        return nil, err
    }
    return td, nil
}
//...
```

### Getting Random image with type and/or tags
You must have at least either type or tags!
FileType consists of: `jpg, png, gif, any`. Jpg and jpeg are treated as equal
NSFW consists of: `false, true, only`
```go
//...
func GetRandomImage(typ string, tags []string,filetype net.FileType,nsfw net.Nsfw, hidden bool) (*data.RandomData, error){
    ri, err := weebgo.GetRandomImage(typ, tags, filetype, nsfw, hidden)
    if err != nil{
        //error handling
        return nil, err
    }
    return ri, nil
}
//...
```

# Models
### TagsData
```go
type TagsData struct{
    status int
    Tags []string
}
```

### TypesData
```go
type TypesData struct{
    status int
    Types []string
}
```

### RandomData
```go
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

type Tags struct{
    Name string
    Hidden bool
    User string
}
```