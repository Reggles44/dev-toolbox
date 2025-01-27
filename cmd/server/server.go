package server

import (
	"dev-server/cmd/server/handlers"
	"dev-server/cmd/server/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	help      bool
	port      uint16
	ServerCmd = &cobra.Command{
		Use:   "server [-p 1234]",
		Short: "Launch a local server",
		Run: func(cmd *cobra.Command, args []string) {
			http.Handle("GET /", middleware.LogMiddleware(handlers.RootGetHandler))
			http.Handle("POST /echo/", middleware.LogMiddleware(handlers.EchoHandler))
			http.Handle("/success/", middleware.LogMiddleware(handlers.SuccessHandler))
			http.Handle("/error/", middleware.LogMiddleware(handlers.ErrorHandler))
			http.Handle("/redirect/", middleware.LogMiddleware(handlers.RedirectHandler))
			http.Handle("/basic_auth/", middleware.LogMiddleware(handlers.BasicAuthHandler))

			log.Printf("Listening on :%d...", port)
			err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
			if err != nil {
				log.Fatalf("Error listening to port (%d): %s", port, err)
			}
		},
	}
)

func init() {
	ServerCmd.Flags().Uint16Var(&port, "port", 8000, "Specify port")
}
