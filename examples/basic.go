package main

import (
	"codeberg.org/redson/gitea-bindings-go/gitea"
	"fmt"
	"log"
)

func main() {
	// The default Gitea instance is Codeberg
	client := gitea.NewClient(nil, nil)

	// redson/sedcord-go
	repo, _ := client.Repositories.Get("redson", "sedcord-go")

	user, err := client.Users.Get("redson")
	if err != nil {
		log.Fatal(err)
	}

	// We're just getting some user infos
	fmt.Printf(
		` === User ===

User: %v
ID:    %v
Name:  %v
Email: %v
Desc:  %v

`, user.Login, user.ID, user.Full_Name, user.Email, user.Description)

	fmt.Printf(
		` === Repositório ===

Full_Name: %v
Stars: %v
Description: %v
Owner: %v

`, repo.FullName, repo.StarsCount, repo.Description, repo.Owner.Login)

}
