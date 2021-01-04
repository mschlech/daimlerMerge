package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Print("getHealthCheck ", w)
	params := mux.Vars(r)
	jsonResponse(w, http.StatusOK, params)
}

func GetReadiness(w http.ResponseWriter, r *http.Request) {
	log.Print("getReadiness", w)
	params := mux.Vars(r)
	jsonResponse(w, http.StatusAccepted, params)
}

func GetMergeList(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var err error
	log.Print(err, vars)
	intervals := []Interval{{25, 30}, {2, 19}, {14, 23}, {4, 8}}
	//intervals := []Interval{{8 ,10}, {2 ,6}, {1 ,3}, {15, 18}}
	//var a = merge1(intervals)
	var start = time.Now().UnixNano()
	var a = merge(intervals)
	var end = time.Now().UnixNano()

	log.Println("interval : ", intervals, "mergedList : ", a, "took :  ", end-start, "nsec")
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
