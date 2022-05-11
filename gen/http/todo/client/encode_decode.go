// Code generated by goa v3.7.3, DO NOT EDIT.
//
// todo HTTP client encoders and decoders
//
// Command:
// $ goa gen todo/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	todo "todo/gen/todo"
	todoviews "todo/gen/todo/views"

	goahttp "goa.design/goa/v3/http"
)

// BuildHelloRequest instantiates a HTTP request object with method and path
// set to call the "todo" service "hello" endpoint
func (c *Client) BuildHelloRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		name string
	)
	{
		p, ok := v.(*todo.HelloPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todo", "hello", "*todo.HelloPayload", v)
		}
		name = p.Name
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: HelloTodoPath(name)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todo", "hello", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeHelloResponse returns a decoder for responses returned by the todo
// hello endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeHelloResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "hello", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todo", "hello", resp.StatusCode, string(body))
		}
	}
}

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "todo" service "show" endpoint
func (c *Client) BuildShowRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*todo.ShowPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todo", "show", "*todo.ShowPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowTodoPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todo", "show", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeShowResponse returns a decoder for responses returned by the todo show
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ShowResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "show", err)
			}
			p := NewShowTodoOK(&body)
			view := "default"
			vres := &todoviews.Todo{Projected: p, View: view}
			if err = todoviews.ValidateTodo(vres); err != nil {
				return nil, goahttp.ErrValidationError("todo", "show", err)
			}
			res := todo.NewTodo(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todo", "show", resp.StatusCode, string(body))
		}
	}
}
