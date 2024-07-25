package ghost

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func (router *Router) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	requestLine, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	split := strings.Split(requestLine, " ")

	handler, httpError := router.findRoute(strings.ToLower(split[0]), strings.TrimSpace(split[1]))

	var response string

	if httpError != nil {
		if httpErrorMessage, ok := httpError.(NotFoundError); ok {
			response = createResponse(404, "Not Found", "text/plain", httpErrorMessage.message)
		} else {
			response = createResponse(500, "Internal Error", "text/plain", "unexpected error")
		}
	} else {
		result := handler()

		response = createResponse(200, "OK", "text/plain", fmt.Sprint(result))
	}

	_, err = conn.Write([]byte(response))

	if err != nil {
		log.Fatal(err.Error())
	}
}

func (router *Router) Listen(port uint) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	log.Printf("Listening at 0.0.0.0:%d\n", port)

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go router.handleConnection(conn)
	}
}
