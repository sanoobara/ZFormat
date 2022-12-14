package main
import (
    "fmt"
    "html/template"
    "net/http"
    "net/url"
    "os"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
    u, err := url.Parse(r.URL.String())
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }

	http.
    params := u.Query()
    searchKey := params.Get("q")
    page := params.Get("page")
    if page == "" {
        page = "1"
    }

    fmt.Println("Search Query is: ", searchKey)
    fmt.Println("Results page is: ", page)
}


var tpl = template.Must(template.ParseFiles("cmd/site/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tpl.Execute(w, nil)
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    mux := http.NewServeMux()

    fs := http.FileServer(http.Dir("assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

  // Add the next line
    mux.HandleFunc("/search", searchHandler)
    mux.HandleFunc("/", indexHandler)
    http.ListenAndServe(":"+port, mux)
}