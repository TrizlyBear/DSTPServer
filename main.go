package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type Server struct {
	Name        string `json:"name"`
	Port        int    `json:"port"`
	Description string `json:"description"`
}

func loadConfig() Server {
	file, err := ioutil.ReadFile("serverconfig.json")
	_ = mux.NewRouter()
	var SConfig Server
	_ = json.Unmarshal([]byte(file), &SConfig)
	if err != nil {
		fmt.Println(err)
	}
	return SConfig
}

func main() {
	config := loadConfig()

	router := mux.NewRouter()
	//router.HandleFunc("/server/{file}",routes.ServerWorld)
	router.PathPrefix("/server/").Handler(http.StripPrefix("/server/", http.FileServer(http.Dir("./server"))))
	http.Handle("/", router)
	http.ListenAndServe("localhost:6969", nil)

	fmt.Println("Hosting server from ", config.Name)
}
