package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//os.Getenv("STORAGE_ROOT") +
func get(w http.ResponseWriter, r *http.Request) {
	f, e := os.Open("./shares/" +
		strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
