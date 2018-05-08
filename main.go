package main

import (
	"fmt"
	"math"
)

func getTotal(slc []int) int {
	ret := 0
	for _, val := range slc {
		ret = ret + val
	}
	return ret
}

func normalFun(ary []int) ([]int, []int) {
	slcA := make([]int, 0)
	slcB := make([]int, 0)

	for _, val := range ary {
		if getTotal(slcA) > getTotal(slcB) {
			slcB = append(slcB, val)
		} else {
			slcA = append(slcA, val)
		}
	}

	fmt.Println(slcA, getTotal(slcA))
	fmt.Println(slcB, getTotal(slcB))
	fmt.Println("")
	return slcA, slcB
}

func diff(ary []int, ary1 []int) int {
	s1 := 0
	for _, val := range ary {
		s1 = s1 + val
	}

	s2 := 0
	for _, val := range ary1 {
		s2 = s2 + val
	}
	return int(math.Abs(float64(s2) - float64(s1)))
}

func doSwap(ary []int, ary1 []int, idx int) {
	tmp := ary[idx]
	ary[idx] = ary1[idx]
	ary1[idx] = tmp
}

func doImFun(ary []int, ary1 []int, swNum int, minDiff int) ([]int, []int, int, bool) {
	i := 0
	bestAry := make([]int, len(ary))
	bestAry1 := make([]int, len(ary))
	bfind := false
	for {
		//swap swNum
		for j := 0; j < swNum; j++ {
			doSwap(ary, ary1, j+i)
		}
		tDiff := diff(ary, ary1)
		if tDiff < minDiff {
			copy(bestAry, ary)
			copy(bestAry1, ary1)
			minDiff = tDiff
			bfind = true
			fmt.Println("bestAry:", bestAry, ", bestAry1:", bestAry1, " diff:", tDiff)
		}
		fmt.Println("ary:", ary, " ary1:", ary1, " swNum:", swNum, " diff:", diff(ary, ary1))
		for j := 0; j < swNum; j++ {
			doSwap(ary, ary1, j+i)
		}
		i += swNum
		if i >= len(ary) || tDiff == 0 {
			break
		}
	}
	if bfind {
		return bestAry, bestAry1, minDiff, true
	} else {
		return nil, nil, 0, false
	}
}

func improveFun(ary []int, ary1 []int) {
	lenv := len(ary) / 2
	minDiff := diff(ary, ary1)
	bary := make([]int, len(ary))
	bary1 := make([]int, len(ary))
	df := 0
	ok := false
	for i := 0; i < lenv; i++ {
		bary, bary1, df, ok = doImFun(ary, ary1, i+1, minDiff)
		if ok {
			minDiff = df
			if df == 0 {
				break
			}
		}
	}
	fmt.Println("finnal result:", bary, ", ", bary1, " diff:", minDiff)
}

func main() {
	ary := []int{18, 16, 15, 14, 8, 7, 3, 1}
	fmt.Println("src array:", ary)
	fmt.Println("normal process")
	fmt.Println("---------------------")
	ary, ary1 := normalFun(ary)

	fmt.Println("start improve process")
	fmt.Println("---------------------")
	improveFun(ary, ary1)
}
