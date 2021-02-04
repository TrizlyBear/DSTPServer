package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ServerWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println(mux.Vars(r)["file"])
	if mux.Vars(r)["file"] == "world" {
		http.ServeFile(w, r, "./server/Floppa.dsv")
	} else {
		fmt.Fprint(w, "Nope")
	}

}
