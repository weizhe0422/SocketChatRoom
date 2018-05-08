package main

import (
	"github.com/googollee/go-socket.io"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			so.BroadcastTo("chat", "chat message", msg)
		})

		so.On("chat ack", func(msg string) string {
			return msg
		})

		so.On("disconnection", func() {
			log.Println("on disconnection")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", c.Handler(server))
	log.Println("Server at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
