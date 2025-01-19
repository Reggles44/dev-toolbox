package main

import (
	"dev-server/cmd/devserver/handlers"
	"dev-server/cmd/devserver/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func runServer(port *uint16) {
	http.Handle("/", middleware.LogMiddleware(handlers.RootHandler))
	http.Handle("/echo/", middleware.LogMiddleware(handlers.EchoHandler))
	http.Handle("/success/", middleware.LogMiddleware(handlers.SuccessHandler))
	http.Handle("/error/", middleware.LogMiddleware(handlers.ErrorHandler))

	log.Printf("Listening on :%d...", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatalf("Error listening to port (%d): %s", *port, err)
	}
}

func main() {
	var help bool
	var port uint16

	serverCmd := &cobra.Command{
		Use:   "dev-server",
		Short: "Launch a local server",
		Run: func(cmd *cobra.Command, args []string) {
			if help {
				cmd.Help()
				return
			}
			runServer(&port)
		},
	}

	serverCmd.Flags().BoolVarP(
		&help,
		"help",
		"h",
		false,
		"Shows command options",
	)

	serverCmd.Flags().Uint16VarP(
		&port,
		"port",
		"p",
		8000,
		"Shows command options",
	)

	serverCmd.Execute()
}
