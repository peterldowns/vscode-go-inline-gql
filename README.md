# vscode-go-inline-gql
Syntax highlighting for inline gql in go.

Raw string literals that start with `#gql\n` will be highlighted as gql.

```go
gql := `#gql
query {
  user {
    login
  }
}
`
```
