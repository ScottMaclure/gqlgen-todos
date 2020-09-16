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

mutation upsertTodo {
  upsertTodo(input:{text:"todo 3", userId:"1"}) {
    user {
      id
    }
    text
    done
  }
}

mutation upsertTodo {
  upsertTodo(input:{id:"T6129484611666145821", text:"todo 3", userId:"1", done: true}) {
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

## TODO

### Sooner

* Add an "update todo" mutation
* Create todos table in postgres
* Fetch todos from postgres
* Save todo to postgres
* Add auth requirement (header)
* Fetch only current User's todos
* Add login feature
* Add create User feature

### Later

* Use Wonka for asynchronous streams of data updates to client
* Subscriptions: Look into WebSockets for getting data (Apollo? urql? Gin support?)

## Changelog

## 2020-09-16

* Move gqlgen-todos into separate github project

## 2020-09-15

* Setup graphql server
* Setup gin-gonic http server
* Setup uhtml client