package ghost

import (
	"testing"
)

func TestAddRouteFunc(t *testing.T) {
	var routes = Router[string]{}

	routes.addRoute("put", "/users/:id/test", func() string { return "" })

	if _, ok := routes["put"]; !ok {
		t.Fatal("Expected to routes have a put property")
	}

	put := routes["put"]

	if put.end {
		t.Fatal("put end property expected to be false")
	}

	if put.handler != nil {
		t.Fatal("put handler property expected to be nil")
	}

	if _, ok := put.next["users"]; !ok {
		t.Fatal("put root tree excepted to have users key")
	}

	put = put.next["users"]

	if put.end {
		t.Fatal("users end property expected to be false")
	}

	if put.handler != nil {
		t.Fatal("users handler property expected to be nil")
	}

	if _, ok := put.next["*"]; !ok {
		t.Fatal("users root tree excepted to have \"*\" key")
	}

	put = put.next["*"]

	if put.end {
		t.Fatal("\"*\" end property expected to be false")
	}

	if put.handler != nil {
		t.Fatal("\"*\" handler property expected to be nil")
	}

	if _, ok := put.next["test"]; !ok {
		t.Fatal("\"*\" root tree excepted to have test key")
	}

	put = put.next["test"]

	if !put.end {
		t.Fatal("test end property expected to be true")
	}

	if put.handler == nil {
		t.Fatal("test handler property expected to be filled")
	}
}

func TestPutWithMoreThanOneRoute(t *testing.T) {
	var routes = Router[string]{}

	routes.addRoute("put", "/users/:id/test", func() string { return "" })
	routes.addRoute("put", "/users/:id/other", func() string { return "" })
	routes.addRoute("put", "/users/other", func() string { return "" })

	route1 := routes["put"].next["users"].next["*"].next["test"]
	route2 := routes["put"].next["users"].next["*"].next["other"]
	route3 := routes["put"].next["users"].next["other"]

	if !route1.end {
		t.Fatal("route1 end expected to be true")
	}

	if !route2.end {
		t.Fatal("route2 end expected to be true")
	}

	if !route3.end {
		t.Fatal("route3 end expected to be true")
	}

	if route1.handler == nil {
		t.Fatal("route1 handler expected to be filled")
	}

	if route2.handler == nil {
		t.Fatal("route2 handler expected to be filled")
	}

	if route3.handler == nil {
		t.Fatal("route3 handler expected to be filled")
	}
}
