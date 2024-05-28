package ghost

import (
	"strings"
)

func CreateRouter() Router {
	return make(Router)
}

func (currentTree *Router) addRoute(method, url string, handler handler) {
	if _, ok := (*currentTree)[method]; !ok {
		(*currentTree)[method] = node{
			end:  false,
			next: make(tree),
		}
	}

	route := splitUrl(url)

	current := (*currentTree)[method]
	isRouteEnd := false

	for i, path := range route {
		isRouteEnd = i == len(route)-1

		var property string

		if strings.HasPrefix(path, ":") {
			property = "*"
		} else {
			property = path
		}

		next := node{
			next: make(map[string]node),
			end:  false,
		}

		if isRouteEnd {
			next.handler = handler
			next.end = true
		}

		if _, ok := current.next[property]; !ok {
			current.next[property] = next
		}

		current = current.next[property]
	}
}

func (currentTree *Router) findRoute(method string, url string) (handler, HttpError) {
	route := splitUrl(url)

	current := (*currentTree)[method]

	for _, path := range route {
		var property string = path

		if _, ok := current.next[property]; !ok {
			if _, starExists := current.next["*"]; !starExists {
				return nil, NotFoundError{
					message: "route not found",
				}
			}

			property = "*"
		}

		current = current.next[property]
	}

	if !current.end {
		return nil, NotFoundError{
			message: "route not found",
		}
	}

	return current.handler, nil
}

func (router Router) Put(url string, handler handler) *Router {
	router.addRoute("put", url, handler)

	return &router
}

func (router Router) Get(url string, handler handler) *Router {
	router.addRoute("get", url, handler)

	return &router
}

func (router Router) Delete(url string, handler handler) *Router {
	router.addRoute("delete", url, handler)

	return &router
}

func (router Router) Patch(url string, handler handler) *Router {
	router.addRoute("patch", url, handler)

	return &router
}

func (router Router) Post(url string, handler handler) *Router {
	router.addRoute("post", url, handler)

	return &router
}

func (router Router) Options(url string, handler handler) *Router {
	router.addRoute("options", url, handler)

	return &router
}

func (router Router) Head(url string, handler handler) *Router {
	router.addRoute("head", url, handler)

	return &router
}
