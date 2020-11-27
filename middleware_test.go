package middlewares

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func registerMiddlewareRandomly(registeredMiddlewares []Middleware) *MiddlewareStack {
	stack := &MiddlewareStack{}
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	sort.Slice(registeredMiddlewares, func(i, j int) bool {
		return r.Intn(100)%2 == 1
	})

	for _, m := range registeredMiddlewares {
		stack.Use(m)
	}

	return stack
}

func TestConflictingMiddlewares(t *testing.T) {
	t.Skipf("conflicting middlewares")
}
