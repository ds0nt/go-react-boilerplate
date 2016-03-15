package routes

import "github.com/ant0ine/go-json-rest/rest"

// CreateMap creates a mindmap
func GetApps(w rest.ResponseWriter, r *rest.Request) {

	w.WriteJson(map[string]string{
		"forky": "http://forky.io",
	})
}
