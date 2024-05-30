package main

import "fmt"

func main() {
	gql := `#gql
query {
  user {
    login
  }
}
	`
	fmt.Println(gql)
}
