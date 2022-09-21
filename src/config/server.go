package config

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/wildanfaz/backendgolang2_week9/src/routers"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start app",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		address := "127.0.0.1:8080"

		if port := os.Getenv("APP_PORT"); port != "" {
			address = "127.0.0.1:" + port
		}

		srv := &http.Server{
			Addr:         address,
			WriteTimeout: time.Second * 20,
			ReadTimeout:  time.Second * 20,
			IdleTimeout:  time.Second * 100,
			Handler:      mainRoute,
		}

		fmt.Println("running on port", os.Getenv("APP_PORT"))
		srv.ListenAndServe()
		return nil
	} else {
		return err
	}
}
