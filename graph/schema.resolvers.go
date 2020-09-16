package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen-todos/graph/generated"
	"gqlgen-todos/graph/model"
	"log"
	"math/rand"
	"strconv"

	"github.com/vektah/gqlparser/gqlerror"
)

func (r *mutationResolver) UpsertTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	// TODO Should work with existing user, not just a random ID.

	done := false
	if input.Done != nil && *input.Done == true {
		done = true
	}

	// TODO This is clumsy. Can an ORM take care of this for me?
	todo := &model.Todo{}
	if input.ID != nil {
		// User is attempting to update existing todo.
		found := false
		for _, v := range r.todos {
			log.Printf("Checking existing ID=%s with new ID=%s", v.ID, *input.ID)
			if v.ID == *input.ID {
				// Existing ID
				todo = v
				found = true
				break
			}
		}
		if !found {
			// See https://gqlgen.com/reference/errors/
			msg := "TODO Can't find todo by ID, how best to fail here?"
			log.Printf(msg)
			return nil, gqlerror.Errorf(msg)
		}
		todo.Text = input.Text
		todo.Done = done
		// Do we need to set it back into the slice? Or pass by reference?
	} else {
		// New TODO from User
		todo.UserID = input.UserID               // set user
		todo.ID = fmt.Sprintf("T%d", rand.Int()) // TODO Leave this to the database autoincrement field.
		todo.Text = input.Text
		todo.Done = done
		r.todos = append(r.todos, todo)
	}

	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {

	r.ConnectDB(ctx)
	// fmt.Printf("Todos, db=%v\n", r.db)

	// pgxscan.Select(ctx, r.db, &r.todos, `select * from todos order by created_at asc`)

	r.todos = nil // empty the slice and underlying data to gc
	rows, err := r.db.Query(ctx, "select id, text, done from todos order by created_at asc")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var text string
		var done bool
		err = rows.Scan(&id, &text, &done)
		if err != nil {
			panic(err)
		}
		fmt.Printf("text=%s, done=%t\n", text, done)
		r.todos = append(r.todos, &model.Todo{
			ID:   strconv.Itoa(id),
			Text: text,
			Done: done,
		})
	}

	return r.todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Username: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
