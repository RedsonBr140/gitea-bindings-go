# gitea-bindings-go
gitea-bindings-go is a Go client library for accessing the [Gitea API v1](https://try.gitea.io/api/swagger).

## Installation
```
go get https://codeberg.org/redson/gitea-bindings-go/gitea
```
You can also import the package
```go
import "https://codeberg.org/redson/gitea-bindings-go"
```
and run `go get` without parameters.

## Usage
```go
import "https://codeberg.org/redson/gitea-bindings-go"
```

Then, you need to construct a new client:
```go
client := gitea.NewClient(nil, nil)

// Get infos about a repository
repo, err := client.Repositories.Get("redson", "gitea-bindings-go")

fmt.Println("Description:", repo.Description)
```

Note that the default gitea instance is Codeberg, if you want to use another, for example:
```go
instanceURL := url.Parse("https://try.gitea.io/api/v1/") // Don't forget the final slash.
client := gitea.NewClient(nil, &gitea.ClientOptions{
  BaseURL: instanceURL,
})
```

## License
This library is distributed under MIT License found in [LICENSE](https://codeberg.org/redson/gitea-bindings-go/src/branch/main/LICENSE) file.