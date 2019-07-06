package main

import (
	"html/template"
	"log"
	"os"
)

type car struct {
	Model   string
	Company string
}

func main() {
	template := template.Must(template.ParseGlob("templates/*"))
	//the last argument is how you pass data you can pass any type of datatype struct array list anything
	err := template.ExecuteTemplate(os.Stdout, "basic.gohtml", car{"Swift", "Maruti"})
	if err != nil {
		log.Fatalln(err)
	}

}
