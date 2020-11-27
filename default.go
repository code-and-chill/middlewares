package middlewares

var DefaultMiddlewareStack = &MiddlewareStack{}

func Use(middleware Middleware) {
	DefaultMiddlewareStack.Use(middleware)
}
