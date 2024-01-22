package greeting

import (
	"context"
	"fmt"

	"encore.dev/beta/errs"
)

// Response is the response data for the Greeting service endpoints.
type Response struct{
	Message string
}

// Params is the request data for the Greeting service endpoints.
type Params struct{
	Name string
}

type Data struct {
    Email string
}


// Ping is an API endpoint that tests the API is online.
//encore:api private method=GET path=/ping
func Ping(ctx context.Context) (*Response, error) {
    msg := fmt.Sprintln("Pong")
    return &Response{Message: msg}, nil
}

// Hello returns a greeting with the name specified
// encore:api public method=GET path=/greet/:name
func Hello(ctx context.Context, name string) (*Response, error)  {
	msg := fmt.Sprintf("Hello, %s!", name)
	return &Response{Message: msg}, nil
}


// Login tests the Authentication
// encore:api auth method=GET path=/auth/login
func Login(ctx context.Context) (*Response,error){
	msg := fmt.Sprintln("Super private endpoint")
	return &Response{Message: msg}, &errs.Error{
        Code: errs.Unauthenticated,
        Message: "invalid token",
    }
}
