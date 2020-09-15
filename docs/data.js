/**
 * DIY GraphQL API Client using browser's Fetch API.
 * TODO What about websockets?
 * TODO What about auth?
 * @see https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API
 */

const QUERY_TODOS = `query findTodos {
    todos {
    text
    done
    user {
      username
    }
  }
}`

const query = (gql) => {
    return fetch('/query', {
        method: 'POST', 
        headers: {'Content-Type': 'application/json', 'Accept': 'application/json'},
        body: JSON.stringify({query: gql})
    }).then(r => {
        if (r.status !== 200) {
            throw new Error(`${r.status} ${r.statusText}`)
        }
        return r.json()
    })
}