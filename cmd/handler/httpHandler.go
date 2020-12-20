package handler

import (
	"fmt"
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
	intervals := []Interval{{25, 30}, {2, 19}, {14, 23}, {4, 8}}
	//intervals := []Interval{{8 ,10}, {2 ,6}, {1 ,3}, {15, 18}}
	//var a = merge1(intervals)
	var a = merge(intervals)
	log.Println("interval : ", intervals, "mergedList : ", a)
	t := map[string]string{"mergedList: ": fmt.Sprintf("%v", a)}
	log.Print("t : ", t)
	jsonResponse(w, http.StatusOK, t)

}

func GetBenchmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	log.Println("params ", params)
	return
}
