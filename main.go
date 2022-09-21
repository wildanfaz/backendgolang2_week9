package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/wildanfaz/backendgolang2_week9/src/routers"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	mainRoute, err := routers.New()

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("APP_PORT")
	fmt.Println("Running On Port", port)

	http.ListenAndServe(port, mainRoute)
}
