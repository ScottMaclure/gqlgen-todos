# gqlgen-todo

## Queries

```graphql

query findTodos {
  	todos {
      text
      done
      user {
        username
      }
    }
}

mutation createTodo {
  createTodo(input:{text:"todo", userId:"1"}) {
    user {
      id
    }
    text
    done
  }
}

query findUsers {
  	users {
      id
      username
    }
}

```

## Using cURL

Note: The playground has a "copy cURL" button :)

```bash
curl 'http://localhost:8080/query' \
-H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' \
-H 'Origin: http://localhost:8080' \
--compressed \
--data-binary '{"query":"query findTodos {\n  \ttodos {\n      text\n      done\n    }\n}"}' 
```

Gives you back:

```json
{"data":{"todos":[{"text":"todo 3","done":false},{"text":"todo 3","done":false},{"text":"todo 4","done":false}]}}
```