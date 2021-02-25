package plan

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.Handle("/state.json", http.FileServer(http.Dir(".")))
	http.Handle("/", http.FileServer(http.Dir("webui")))

	fmt.Println("serving at http://localhost:8080")
	panic(http.ListenAndServe(":8080", nil))
}
