package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/junlyang/lyangcoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml"))
	data := homeData{"안녕", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}
func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("second block")

	http.HandleFunc("/", home)
	fmt.Printf("Lisnening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
