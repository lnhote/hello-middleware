package main

import (
	"fmt"
)

type Handler func()

type MiddleWare func(next Handler) Handler

var middlewares = make([]MiddleWare, 0)

/**
The middleware is FIFO
*/
func Use(middleware MiddleWare) {
	middlewares = append(middlewares, middleware)
}

func ProcessRequest(handler Handler) {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	handler()
}

func loggerMiddleware(next Handler) Handler {
	return func() {
		fmt.Println("[loggerMiddleware]Log: Get a request.")
		next()
		fmt.Println("[loggerMiddleware]Log: You just processed a request.")
	}
}
func authMiddleware(next Handler) Handler {
	return func() {
		fmt.Println("[authMiddleware]Before handling the request, let's authenticate the user...")
		fmt.Println("[authMiddleware]You are good man.")
		next()
	}
}
func dbMiddleware(next Handler) Handler {
	return func() {
		fmt.Println("[dbMiddleware]Connect to the database.")
		next()
		fmt.Println("[dbMiddleware]Disconnect from database.")
	}
}

func jsonMiddleware(next Handler) Handler {
	return func() {
		next()
		fmt.Println("[jsonMiddleware]Format the data to json for HTTP response.")
	}
}

func panicMiddleware(next Handler) Handler {
	return func() {
		next()
		fmt.Println("[panicMiddleware]Deal with panic, or you will be fucked up!")
	}
}

func formMiddleware(next Handler) Handler {
	return func() {
		fmt.Println("[formMiddleware]Normalize the post form/get params.")
		next()
	}
}

func mainRequestHandler() {
	fmt.Println("Hello world")
}

func main() {

	//-------------------
	// Custom middleware, FIFO - mainHandler - LIFO
	//-------------------
	Use(panicMiddleware)
	Use(loggerMiddleware)
	Use(formMiddleware)
	Use(authMiddleware)
	Use(jsonMiddleware)
	Use(dbMiddleware)

	// Run the main handler with middlewares
	ProcessRequest(mainRequestHandler)
}
