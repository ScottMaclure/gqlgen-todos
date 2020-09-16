const {log} = console;

const DATA_STATES = {
    DEFAULT: 0,
    LOADING: 1,
    COMPLETE: 2,
    ERROR: 3
}

const globalState = {
    version: 'BETA 001',
    button: {
        disabled: false, 
        text: 'Click Me'
    },
    todoList: {
        comment: "GraphQL backed.",
        state: DATA_STATES.DEFAULT,
        todos: []
    }
};

const handleTestClick = () => {
    globalState.button.text = `Clicked at ${Date.now()}`
    renderApp()
}

const loadTodos = () => {
    let {todoList} = globalState
    todoList.state = DATA_STATES.LOADING
    renderApp()

    query(QUERY_TODOS).then(data => {
        todoList.state = DATA_STATES.COMPLETE
        globalState.todoList.todos = data.data.todos
    }).catch(err => {
        log(`loadTodos err=${err}`)
        todoList.state = DATA_STATES.ERROR
        todoList.todos = []
    }).finally(() => {
        renderApp()
    })
}

const mainElement = document.getElementsByTagName('main')[0]

// (re)render the entire app
const renderApp = () => {
    uhtml.render(mainElement, uhtml.html`
        <header>
            <h1>GraphQL Client Test</h1>
            <small>Version: ${globalState.version}</small>
        </header>
        <h2>Test Button</h2>
        <div>
            ${Button(globalState.button, handleTestClick)}
            ${Button({text: 'Disabled Button', disabled: true}, handleTestClick)}
            ${Button({text: 'Load Todos'}, loadTodos)}
        </div>
        <h2>Todo List</h2>
        <div>
            ${TodoList(globalState.todoList)}
        </div>
    `);
}

const Button = (buttonState, clickHandler) => uhtml.html`
    <button class="clickable" onclick=${clickHandler} .disabled=${buttonState.disabled||false}>
        ${buttonState.text}
    </button>
`

const TodoList = (todoList) => {
    if (todoList.state == DATA_STATES.DEFAULT) {
        return uhtml.html`<p>Press the button to load todos.</p>`
    // } else if (todoList.state == DATA_STATES.LOADING) {
        // Disable loading indicator to test out uhtml DOM updates.
        // return uhtml.html`<p>Loading...</p>`
    } else if (todoList.state == DATA_STATES.ERROR) {
        return uhtml.html`<p>Error loading todo list, check logs.</p>`
    } else if (todoList.state == DATA_STATES.LOADING || todoList.state == DATA_STATES.COMPLETE) {
        if (todoList.todos.length > 0) {
            // TODO Use keyed rendering?
            return uhtml.html`
            <ul>
                ${todoList.todos.map(
                    (item, i) => {
                        let cssStyle = item.done ? 'text-decoration: line-through;' : ''
                        return uhtml.html`<li style="${cssStyle}">${i + 1}: ${item.text}</li>`
                    }
                )}
            </ul>
            `
        } else {
            return uhtml.html`<p>No todos available. Create one?</p>`
        }
    }
}

// Kick the whole thing off on load.
renderApp()
