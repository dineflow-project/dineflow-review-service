package main

import (
	"log"
	"net/http"
	"runtime"

	"dineflow-review-service/configs"
	"dineflow-review-service/routes"
	"dineflow-review-service/utils/inits"

	"github.com/gorilla/mux"
)

func main() {
	inits.InitializeDatabase()

	router := mux.NewRouter()
	routes.ProtectedRoute(router)

	local_os := runtime.GOOS
	if local_os == "windows" {
		log.Fatal(http.ListenAndServe("127.0.0.1:"+configs.EnvServicePort(), nil))
	} else {
		log.Fatal(http.ListenAndServe(":"+configs.EnvServicePort(), nil))
	}
}
