package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"dineflow-review-services/routes"
	"dineflow-review-services/utils/inits"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	inits.InitializeDatabase()

	router := mux.NewRouter()
	routes.ProtectedRoute(router)

	local_os := runtime.GOOS
	if local_os == "windows" {
		log.Fatal(http.ListenAndServe("127.0.0.1:"+os.Getenv("PORT"), nil))
	} else {
		log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	}
}
