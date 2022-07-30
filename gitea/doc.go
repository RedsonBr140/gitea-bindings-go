/*
Package gitea provides Go bindings for the Gitea API.

Usage

Construct a new Client, and then you can use the services
to access different parts of the API. For example:
    client := gitea.NewClient(nil)
    
    // get info about a repository
    repo, err := client.Repositories.Get("redson", "gitea-bindings-go")

    fmt.Println("Description:", repo.Description)

The default instance is Codeberg, if you want to use another, do something like:
  instanceURL := url.Parse("https://gittea.dev/api/v1/")
  client := gitea.NewClient(&gitea.ClientOptions{
    baseURL: instanceURL,
  })

Authentication

Authentication is the same as changing the instance
but now what you provide is a token, for example:
  client := gitea.NewClient(&gitea.ClientOptions{
    Token: "foo",
  })

For more info, see the examples at https://codeberg.org/redson/gitea-bindings-go/src/branch/main/example.

*/
package gitea // import "https://codeberg.org/redson/gitea-bindings-go/gitea"
