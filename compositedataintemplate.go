package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	template := template.Must(template.ParseGlob("templates/*"))
	cars := []string{"aston martin", "mini cooper", "porsche", "lamborghini"}
	alphabets := map[string]string{
		"a": "apple",
		"b": "ball",
	}
	//the last argument is how you pass data you can pass any type of datatype struct array list anything
	err := template.ExecuteTemplate(os.Stdout, "compositestringslice.gohtml", cars)
	err = template.ExecuteTemplate(os.Stdout, "compositemap.gohtml", alphabets)
	if err != nil {
		log.Fatalln(err)
	}

}
