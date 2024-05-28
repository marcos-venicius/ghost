package ghost

type handler func() string

type node struct {
	next    tree
	handler handler
	end     bool
}

type tree map[string]node

type Router tree
