# gitea-bindings-go
gitea-bindings-go is a Go client library for accessing the [Gitea API v1](https://try.gitea.io/api/swagger).

## Installation
```
go get codeberg.org/redson/gitea-bindings-go/gitea
```
You can also import the package
```go
import "codeberg.org/redson/gitea-bindings-go/gitea"
```
and run `go get` without parameters.

## Usage
```go
import "codeberg.org/redson/gitea-bindings-go/gitea"
```

Construct a new Client, and then you can use the services
to access different parts of the API. For example:
```go
client := gitea.NewClient(nil)
    
// get info about a repository
repo, err := client.Repositories.Get("redson", "gitea-bindings-go")
fmt.Println("Description:", repo.Description)
```
The default instance is Codeberg, if you want to use another, do something like:
```go
instanceURL := url.Parse("https://gittea.dev/api/v1/")
client := gitea.NewClient(&gitea.ClientOptions{
  baseURL: instanceURL,
})
```

### Authentication

Authentication is the same as changing the instance
but now what you provide is a token, for example:
```go
client := gitea.NewClient(&gitea.ClientOptions{
  Token: "foo",
})
```

For more info, see the [examples](https://codeberg.org/redson/gitea-bindings-go/src/branch/main/examples).


## License
This library is distributed under MIT License found in [LICENSE](https://codeberg.org/redson/gitea-bindings-go/src/branch/main/LICENSE) file.

