// This program demonstrates how to use a token with godotenv
package main

import (
	"fmt"
	"log"
	"os"

	"codeberg.org/redson/gitea-bindings-go/gitea"
	"github.com/joho/godotenv"
)

func main() {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Cannot load .env file")
  }

  token := os.Getenv("TOKEN")

  client := gitea.NewClient(&gitea.ClientOptions{
    Token: token,
  })
  user, err := client.Users.Get("")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Logged in as %v (ID: %v)\n", user.Username, user.ID)
}
