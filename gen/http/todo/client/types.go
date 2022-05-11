// Code generated by goa v3.7.3, DO NOT EDIT.
//
// todo HTTP client types
//
// Command:
// $ goa gen todo/design

package client

import (
	todoviews "todo/gen/todo/views"
)

// ShowResponseBody is the type of the "todo" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// IsDone
	IsDone *bool `form:"is_done,omitempty" json:"is_done,omitempty" xml:"is_done,omitempty"`
}

// NewShowTodoOK builds a "todo" service "show" endpoint result from a HTTP
// "OK" response.
func NewShowTodoOK(body *ShowResponseBody) *todoviews.TodoView {
	v := &todoviews.TodoView{
		ID:     body.ID,
		Title:  body.Title,
		IsDone: body.IsDone,
	}

	return v
}