package main

import (
	"bufio"
	"log"
	"net/http"
	"os"

	"./objects"
	"./stateinf"
)

//os.Getenv("LISTEN_ADDRESS")
func main() {
	file, _ := os.Open("./parameter.txt")
	reader := bufio.NewReader(file)
	str, _ := reader.ReadString('\n')
	http.HandleFunc("/shares/", objects.Handler)
	http.HandleFunc("/sizequery/", stateinf.Handler)
	log.Fatal(http.ListenAndServe(str, nil))
}
