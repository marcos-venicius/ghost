package ghost

type RequestHandler func() interface{}

type node struct {
	next    tree
	handler RequestHandler
	end     bool
}

type tree map[string]node

type Router tree
