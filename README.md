# [ent](https://entgo.io/) - reproduce

This repo connectes to this [issue](https://github.com/ent/ent/issues/2838)

## The issue

> interface conversion: interface {} is string, not \*entgen.Cursor

## Structure

- ent:<br/>
  --schema: all ent schemas<br/>
  --rules: all ent privacy rules

- entgen: ent generated code
- tests: golang tests for reproducing the issue

## How to reproduce

1. Connet to local MySQL via `root:root@tcp(localhost:3306)/enttry?parseTime=true`, please update it in main.go.
1. `go run ./main.go`
1. send a request to `http://localhost:8081/create` for creating the mock data
1. then try the below 2 queries

> Please update the ID to the data in your database

**This one works**

```graphql
query user {
  node(id: "user_2D8MrI9P39N18Xq5RBc0F45HYVw") {
    ... on User {
      id
      profiles {
        edges {
          node {
            id
            ownerID
            tenantID
          }
        }
      }
    }
  }
}
```

**This one NOT works**

```graphql
query user {
  node(id: "user_2D8MrI9P39N18Xq5RBc0F45HYVw") {
    ... on User {
      id
      profiles(first: 50, after: "profile_2D8MrHIsrHobAOCXaAtoF13ahIJ") {
        edges {
          node {
            id
            ownerID
            tenantID
          }
        }
      }
    }
  }
}
```
