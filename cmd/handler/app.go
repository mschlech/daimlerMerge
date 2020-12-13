package handler

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
)

/* the struct for the beginning and the end of the interval e.g
{1,3} {2,5} the interval starts at 1 ends at 5
*/
type Interval struct {
	offset, end int
}

/**
merge gets the interval of number , unsorted tuples, but sorting first
O(n log n) problem sorting , n is the number of elements to be sorted
*/
func merge1(intervals []Interval) []Interval{
	fmt.Println("Interval : ", intervals)
	if len(intervals) == 0 {
		log.Print("given interval empty")
	}

	//It sorts a slice using a provided function
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		log.Println("a offset :", a.offset, "b offset :", b.offset)

		log.Println("a end :", a.end, "b end :", b.end)
		log.Println("t : ", a.offset < b.offset || (a.offset == b.offset && a.end > b.end))
		b2 := a.offset < b.offset || (a.offset == b.offset && a.end > b.end)

		return b2

	})

	//To sort the slice while keeping the original order of equal elements
	//sort.SliceStable(intervals, func(i, j int) bool {
	//	a, b := intervals[i], intervals[j]
	//	fmt.Println("a:", a.offset)
	//	fmt.Println("b:", b.offset)
	//	return a.offset < b.offset || (a.offset == b.offset && a.end > b.end)
	//})

	res := []Interval{}
	cur := intervals[0]

	cur, res = getInterval(intervals, cur, res)
	res = append(res, cur)
	//fmt.Println("result : ", res)
	return res
}

func getInterval(intervals []Interval, cur Interval, res []Interval) (Interval, []Interval) {
	//fmt.Println("given interval : ", intervals, "cur : ", cur)
	for _, p := range intervals {
		fmt.Println("p.offset : ", p.offset, "cur end: ", cur.end, "p : ", p)
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

func generate(randMatrix [][]int, givenLength int)  {


	//var v = []Interval{}
	tmp:=make([]Interval,givenLength)
	var a []Interval
	for i := range tmp {
		for j := range tmp {
		tmp[j].offset = rand.Intn(givenLength)
		tmp[i].end= rand.Intn(givenLength)
			log.Println("i:",tmp[i], "j:" , tmp[j])
			//a = Interval{tmp[i].offset,tmp[j].end}
		}

	}

	log.Println(a)
	////var randomOut [][]int
	//for i, innerArray := range randMatrix {
	//	for j := range innerArray {
	//		randMatrix[i][j] = rand.Intn(givenLength)
	//	}
	//}
	//fmt.Println("randomMatrix :"  ,randMatrix)
	//return v
}

//
//func main() {
//	//result :  &[{2 19} {25 30}]
//	intervals := []Interval{{25, 30}, {2, 19}, {14, 23}, {4, 8}}
//    //A := []Interval{}
//	rand.Seed(time.Now().Unix())
//	//	intervals2 := []Interval{generate()}
//	//intervalsRandom := []Interval{rand.Perm(100), rand.Perm(100)}
//	merge1(intervals)
//	//merge1(intervalsRandom)
//	var givenLength = 100
//	randMatrix := make([][]int, givenLength)
//	for i := 0; i < 100; i++ {
//		randMatrix[i] = make([]int, 2)
//		//fmt.Println("out: ", randMatrix[i])
//		//append(A, randMatrix[i])
//	}
//	generate(randMatrix,givenLength)
//
//	fmt.Println(randMatrix)
//
//}
