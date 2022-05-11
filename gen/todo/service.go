// Code generated by goa v3.7.3, DO NOT EDIT.
//
// todo service
//
// Command:
// $ goa gen todo/design

package todo

import (
	"context"
	todoviews "todo/gen/todo/views"
)

// Service that manage todo.
type Service interface {
	// Hello implements hello.
	Hello(context.Context, *HelloPayload) (res string, err error)
	// Show implements show.
	Show(context.Context, *ShowPayload) (res *Todo, err error)
	// Create implements create.
	Create(context.Context, *CreatePayload) (res string, err error)
	// Update implements update.
	Update(context.Context, *UpdatePayload) (res string, err error)
	// Delete implements delete.
	Delete(context.Context, *DeletePayload) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "todo"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"hello", "show", "create", "update", "delete"}

// CreatePayload is the payload type of the todo service create method.
type CreatePayload struct {
	// Title
	Title string
}

// DeletePayload is the payload type of the todo service delete method.
type DeletePayload struct {
	// ID
	ID int
}

// HelloPayload is the payload type of the todo service hello method.
type HelloPayload struct {
	// Name
	Name string
}

// ShowPayload is the payload type of the todo service show method.
type ShowPayload struct {
	// ID
	ID int
}

// Todo is the result type of the todo service show method.
type Todo struct {
	// ID
	ID *int
	// Title
	Title *string
	// IsDone
	IsDone *bool
}

// UpdatePayload is the payload type of the todo service update method.
type UpdatePayload struct {
	// ID
	ID int
	// IsDone
	IsDone bool
}

// NewTodo initializes result type Todo from viewed result type Todo.
func NewTodo(vres *todoviews.Todo) *Todo {
	return newTodo(vres.Projected)
}

// NewViewedTodo initializes viewed result type Todo from result type Todo
// using the given view.
func NewViewedTodo(res *Todo, view string) *todoviews.Todo {
	p := newTodoView(res)
	return &todoviews.Todo{Projected: p, View: "default"}
}

// newTodo converts projected type Todo to service type Todo.
func newTodo(vres *todoviews.TodoView) *Todo {
	res := &Todo{
		ID:     vres.ID,
		Title:  vres.Title,
		IsDone: vres.IsDone,
	}
	return res
}

// newTodoView projects result type Todo to projected type TodoView using the
// "default" view.
func newTodoView(res *Todo) *todoviews.TodoView {
	vres := &todoviews.TodoView{
		ID:     res.ID,
		Title:  res.Title,
		IsDone: res.IsDone,
	}
	return vres
}
