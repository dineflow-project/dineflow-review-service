package main

import (
	"log"
	"net/http"
	"runtime"

	"dineflow-review-services/configs"
	"dineflow-review-services/routes"
	"dineflow-review-services/utils/inits"

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
