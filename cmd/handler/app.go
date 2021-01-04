package handler

import (
	"fmt"
	"log"
	"sort"
)

/**
the app can also run on its own would only need a main method
*/
import (
	"math/rand"
	"time"
)

/* the struct for the beginning and the end of the interval e.g
{1,3} {2,5} the interval starts at 1 ends at 5
*/
type Interval struct {
	offset, end int
}

/**
merge1 as an alternative approach. Gets the interval of number , unsorted tuples, but sorting first
O(n log n) problem sorting , n is the number of elements to be sorted
*/
func merge1(intervals []Interval) []Interval {
	fmt.Println("Interval : ", intervals)
	if len(intervals) == 0 {
		log.Print("given interval empty")
	}

	//To sort the slice while keeping the original order of equal elements
	sort.SliceStable(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		fmt.Println("a:", a.offset)
		fmt.Println("b:", b.offset)
		return a.offset < b.offset || (a.offset == b.offset && a.end > b.end)
	})

	res := []Interval{}
	cur := intervals[0]

	cur, res = getInterval(intervals, cur, res)
	res = append(res, cur)
	//fmt.Println("result : ", res)
	return res
}

/**
method which is invoked only by the alternative merge1 method.
See above. might have one issue as it was missing an interval result
*/
func getInterval(intervals []Interval, cur Interval, res []Interval) (Interval, []Interval) {
	log.Print("given interval : ", intervals, "cur : ", cur)
	for _, p := range intervals {
		log.Print("p.offset : ", p.offset, "cur end: ", cur.end, "p : ", p)
		if p.offset > cur.end {
			res = append(res, cur)
			cur = p
			log.Println("res : ", res, "p: ", p)
		} else {
			if p.offset > cur.end {
				cur.end = p.end
			}
		}
	}
	log.Println("cur ", cur, "res", res)
	return cur, res
}

/**
function which merge the intervals to a result interface of type  interval.
*/
func merge(ivs []Interval) []Interval {
	m := append([]Interval(nil), ivs...)
	if len(m) <= 1 {
		return m
	}
	// just for curiosity. Could be exposed to a logging interface to provide this to an external logging entity.
	// checking on thresholds could lead to an autoscaling approach.
	var startBoolCheck = time.Now().UnixNano()
	checkInRange(m)
	var endBoolCheck = time.Now().UnixNano()

	log.Print("check funcName took: ", endBoolCheck-startBoolCheck, "nsec")

	j := 0
	for i := 1; i < len(m); i++ {
		if m[j].end >= m[i].offset {
			if m[j].end < m[i].end {
				m[j].end = m[i].end
			}
		} else {
			j++
			m[j] = m[i]
		}

	}
	return append([]Interval(nil), m[:j+1]...)
}

/**
implcitly private method which checks if the range within a touple fits and returns a boolean
*/
func checkInRange(m []Interval) {
	sort.Slice(m,
		func(i, j int) bool {
			if m[i].offset < m[j].offset {
				return true
			}
			if m[i].end == m[j].offset && m[i].end < m[j].end {
				return true
			}
			return false
		},
	)
}

/**
not in use as of time constraints. Method should provide a random list of interface intervals.
*/
func generate(randMatrix [][]int, givenLength int) {

	//var v = []Interval{}
	tmp := make([]Interval, givenLength)
	var a []Interval
	for i := range tmp {
		for j := range tmp {
			tmp[j].offset = rand.Intn(givenLength)
			tmp[i].end = rand.Intn(givenLength)
			log.Println("i:", tmp[i], "j:", tmp[j])
			//a = Interval{tmp[i].offset,tmp[j].end}
		}

	}
	log.Println(a)
}
