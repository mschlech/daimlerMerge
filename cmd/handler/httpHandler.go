package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

// globals
var durationmerge int64

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
	var status string
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
		getMergeLatency.WithLabelValues(status).Observe(v)
	}))
	defer func() {
		timer.ObserveDuration()
	}()

	vars := mux.Vars(r)
	var err error
	log.Print(err, vars)
	intervals := []Interval{{25, 30}, {2, 19}, {14, 23}, {4, 8}}
	//intervals := []Interval{{8 ,10}, {2 ,6}, {1 ,3}, {15, 18}}
	//var a = merge1(intervals)
	var start = time.Now().UnixNano()
	var a = merge(intervals)
	var end = time.Now().UnixNano()
	durationmerge = end - start
	log.Println("interval : ", intervals, "mergedList : ", a, "took :  ", end-start, "nsec")
	t := map[string]string{"mergedList: ": fmt.Sprintf("%v", a)}
	log.Print("mergeList  : ", t, " in ", durationmerge, " nsec.")
	jsonResponse(w, http.StatusOK, t)

}

// non functional section

// prometheus specifics
var getMergeLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_request_get_metrics_of_merge_duration_seconds",
	Help:    "get the latency of a merge operation merging intervals",
	Buckets: prometheus.LinearBuckets(0.01, 0.05, 10),
},
	[]string{"status"},
)

// initialize and register the getMergeLatency method
func init() {
	log.Print("register prom getMergeLatency")
	prometheus.MustRegister(getMergeLatency)
}

// simple monitoring approach without premetheus
func GetMetrics(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	log.Println("params ", params)
	metric := map[string]string{"duration of last operation : ": fmt.Sprint("#{a}")}
	jsonResponse(w, http.StatusOK, metric)
}
