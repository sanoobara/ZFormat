package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var tpl = template.Must(template.ParseFiles("cmd/site/index1.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func Timer(times string) {
	times = time.Now().Format("15:04:05\n")
	time.Sleep(1 * time.Second)
	fmt.Println(times)
}

func main() {
	for {
		var timer string
		go Timer(timer)

		//conf, er := config.BulkParser("config\\bulkconfig.json")

		// if er != nil {
		// 	fmt.Println(er)
		// }
		 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			tmpl, _ := template.ParseFiles("cmd/site/index1.html")
			tmpl.Execute(w, timer)
		})

		// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// 	tmpl, _ := template.ParseFiles("cmd/site/index.html")
		// 	tmpl.Execute(w, conf)
		// })

		 fmt.Println("Server is listening...")

		 http.ListenAndServe(":8182", nil)
		 time.Sleep(1 * time.Second)
	}
}
