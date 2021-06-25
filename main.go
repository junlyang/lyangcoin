package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/junlyang/lyangcoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"안녕", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)

	}
	templates.ExecuteTemplate(rw, "add", nil)
}

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("second block")

	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)

	fmt.Printf("Lisnening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
