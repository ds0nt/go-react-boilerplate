package server

import (
	"log"

	"ds0nt.com/config"
	"ds0nt.com/server/routes"

	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

// Run forky
func Run() {
	log.Println("Starting ds0nt.com server")
	go fileServer(config.All.WebRoot, config.All.FileServerPort)
	apiServer(config.All.APIPort)
}

func apiServer(addr string) {
	// Api
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true // origin == "http://127.0.0.1"
		},
		AllowedMethods:                []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:                []string{"Accept", "Content-Type", "Authorization", "Origin"},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})

	router, err := rest.MakeRouter(
		rest.Get("/apps", routes.GetApps),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)

	log.Printf("Api Server Listening on %s\n", addr)
	http.ListenAndServe(addr, api.MakeHandler())
}

func fileServer(path string, addr string) {
	// Static Files
	fs := http.FileServer(http.Dir(path))
	http.Handle("/", fs)

	log.Println("File Server Listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
