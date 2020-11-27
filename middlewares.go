package middlewares

type MiddlewareStack struct {
	middlewares []*Middleware
}

func (stack *MiddlewareStack) Use(middleware Middleware) {
	stack.middlewares = append(stack.middlewares, &middleware)
}

func (stack *MiddlewareStack) Remove(name string) {
	registeredMiddlewares := stack.middlewares
	for idx, middleware := range registeredMiddlewares {
		if middleware.Name == name {
			if idx > 0 {
				stack.middlewares = stack.middlewares[0 : idx-1]
			} else {
				stack.middlewares = []*Middleware{}
			}

			if idx < len(registeredMiddlewares)-1 {
				stack.middlewares = append(stack.middlewares, registeredMiddlewares[idx+1:]...)
			}
		}
	}
}
