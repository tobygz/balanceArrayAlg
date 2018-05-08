package main

import (
        "fmt"
        "math"
)

func diff(ary []uint32, ary1 []uint32) uint32 {
        s1 := uint32(0)
        for _, val := range ary {
                s1 = s1 + val
        }

        s2 := uint32(0)
        for _, val := range ary1 {
                s2 = s2 + val
        }
        return uint32(math.Abs(float64(s2) - float64(s1)))
}

func doSwap(ary []uint32, ary1 []uint32, idx uint32) {
        tmp := ary[idx]
        ary[idx] = ary1[idx]
        ary1[idx] = tmp
}

func doImFun(ary []uint32, ary1 []uint32, swNum uint32, minDiff uint32) ([]uint32, []uint32, uint32, bool) {
        i := uint32(0)
        bestAry := make([]uint32, len(ary))
        bestAry1 := make([]uint32, len(ary))
        bfind := false
        for {
                //swap swNum
                for j := uint32(0); j < swNum; j++ {
                        doSwap(ary, ary1, j+i)
                }
                tDiff := diff(ary, ary1)
                if tDiff < minDiff {
                        copy(bestAry, ary)
                        copy(bestAry1, ary1)
                        minDiff = tDiff
                        bfind = true
                }
                for j := uint32(0); j < swNum; j++ {
                        doSwap(ary, ary1, j+i)
                }
                i += swNum
                if i >= uint32(len(ary)) || tDiff == 0 {
                        break
                }
        }
        if bfind {
                return bestAry, bestAry1, minDiff, true
        } else {
                return nil, nil, 0, false
        }
}

func improvedBalFun(ary []uint32, ary1 []uint32) ([]uint32, []uint32) {
        lenv := uint32(len(ary) / 2)
        minDiff := diff(ary, ary1)
        bary := make([]uint32, len(ary))
        bary1 := make([]uint32, len(ary))
        df := uint32(0)
        ok := false
        for i := uint32(0); i < lenv; i++ {
                bary, bary1, df, ok = doImFun(ary, ary1, i+1, minDiff)
                if ok {
                        minDiff = df
                        if df == 0 {
                                break
                        }
                }
        }
        if !ok {
                return ary, ary1
        }else{
                return bary, bary1
        }
}

func getTotal(slc []uint32) uint32 {
        ret := uint32(0)
        for _, val := range slc {
                ret = ret + val
        }
        return ret
}

func normalFun(ary []uint32) ([]uint32, []uint32) {
        slcA := make([]uint32, 0)
        slcB := make([]uint32, 0)
        limitLen := len(ary) / 2
        for _, val := range ary {
                if getTotal(slcA) > getTotal(slcB)  && len(slcB) < limitLen{
                        slcB = append(slcB, val)
                } else {
                        slcA = append(slcA, val)
                }
        }
        return slcA, slcB
}

func BestBalFun(ary []uint32) ([]uint32, []uint32) {
        ary, ary1 := normalFun(ary)
        ary, ary1 = improvedBalFun(ary, ary1)
        return ary, ary1
}

func main() {
        ary := []uint32{18, 16, 15, 14, 8, 7, 3, 1}
        fmt.Println(BestBalFun(ary))
}
