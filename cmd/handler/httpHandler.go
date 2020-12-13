package handler
import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Print("getHealthCheck ", w)
	params := mux.Vars(r)
	jsonResponse(w, http.StatusOK, params)
}


func GetMergeList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	log.Print(err, vars)

	return
}

func GetBenchmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","application/json")
	params := mux.Vars(r)
	log.Println("params " , params)
	return
}

