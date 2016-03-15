package client

import (
	"log"
	"os"
	"text/template"

	"ds0nt.com/config"
)

func RenderIndex() error {
	log.Printf("Rendering %s -> %s", config.All.IndexTemplate, config.All.IndexTarget)
	file, err := os.Create(config.All.IndexTarget)
	if err != nil {
		log.Println("Template CreateFile Error:", err)
		return err
	}
	defer file.Close()

	t := template.Must(template.ParseFiles(config.All.IndexTemplate))
	err = t.Execute(file, config.All)
	if err != nil {
		log.Println("Template Error:", err)
		return err
	}

	return nil
}
