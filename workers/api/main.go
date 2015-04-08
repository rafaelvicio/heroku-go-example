package main

import (
	"example/workers/api/controllers"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.GetHome)

	n := negroni.New()
	n.UseHandler(router)

	port := os.Getenv("PORT")
	n.Run(":" + port)
}
