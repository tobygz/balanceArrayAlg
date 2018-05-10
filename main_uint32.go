package main

import (
        "fmt"
        "math"
)

type Player struct {
        score uint32
        id    uint32
}

func (this *Player) GetStarForMatch() uint32 {
        return this.score
}

func diff(ary []*Player, ary1 []*Player) uint32 {
        s1 := uint32(0)
        for _, p := range ary {
                s1 = s1 + p.GetStarForMatch()
        }

        s2 := uint32(0)
        for _, p := range ary1 {
                s2 = s2 + p.GetStarForMatch()
        }
        return uint32(math.Abs(float64(s2) - float64(s1)))
}

func doSwap(ary []*Player, ary1 []*Player, idx uint32) {
        ary[idx], ary1[idx] = ary1[idx], ary[idx]
}

func doImFun(ary []*Player, ary1 []*Player, swNum uint32, minDiff uint32) ([]*Player, []*Player, uint32, bool) {
        i := uint32(0)
        bestAry := make([]*Player, len(ary))
        bestAry1 := make([]*Player, len(ary))
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

func improvedBalFun(ary []*Player, ary1 []*Player) ([]*Player, []*Player) {
        lenv := uint32(len(ary) / 2)
        minDiff := diff(ary, ary1)
        bary := make([]*Player, len(ary))
        bary1 := make([]*Player, len(ary))
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
        } else {
                return bary, bary1
        }
}

func getTotal(slc []*Player) uint32 {
        ret := uint32(0)
        for _, p := range slc {
                ret = ret + p.GetStarForMatch()
        }
        return ret
}

func normalFun(ary []*Player) ([]*Player, []*Player) {
        slcA := make([]*Player, 0)
        slcB := make([]*Player, 0)

        limitLen := len(ary) / 2
        for _, p := range ary {
                if getTotal(slcA) > getTotal(slcB) {
                        if len(slcB) < limitLen {
                                slcB = append(slcB, p)
                        } else {
                                slcA = append(slcA, p)
                        }
                } else {
                        if len(slcA) < limitLen {
                                slcA = append(slcA, p)
                        } else {
                                slcB = append(slcB, p)
                        }
                }
        }
        return slcA, slcB
}

func BestBalFun(ary []*Player) ([]*Player, []*Player) {
        debugStr := ""
        for _, p := range ary {
                debugStr = fmt.Sprintf("%s %d", debugStr, p.GetStarForMatch())
        }
        ary1, ary2 := normalFun(ary)
        debugStr = fmt.Sprintf("%s; normalFun:[", debugStr)
        for _, p := range ary1 {
                debugStr = fmt.Sprintf("%s %d", debugStr, p.GetStarForMatch())
        }
        debugStr = fmt.Sprintf("%s; ", debugStr)
        for _, p := range ary2 {
                debugStr = fmt.Sprintf("%s %d", debugStr, p.GetStarForMatch())
        }
        debugStr = fmt.Sprintf("%s diff:%d] improveFun:[", debugStr, diff(ary1, ary2))

        //fmt.Println(ary1, ary2)
        ary1, ary2 = improvedBalFun(ary1, ary2)

        for _, p := range ary1 {
                debugStr = fmt.Sprintf("%s %d", debugStr, p.GetStarForMatch())
        }
        debugStr = fmt.Sprintf("%s; ", debugStr)
        for _, p := range ary2 {
                debugStr = fmt.Sprintf("%s %d", debugStr, p.GetStarForMatch())
        }
        debugStr = fmt.Sprintf("%s diff:%d]", debugStr, diff(ary1, ary2))
        fmt.Println(debugStr)
        return ary1, ary2
}

func main() {
        /*
                ary := []*Player{&Player{score: 18, id: 0},
                        &Player{score: 16, id: 1},
                        &Player{score: 15, id: 2},
                        &Player{score: 14, id: 3},
                        &Player{score: 8, id: 4},
                        &Player{score: 7, id: 5},
                        &Player{score: 3, id: 6},
                        &Player{score: 1, id: 7}}
        */
        ary := []*Player{
                &Player{score: 18, id: 0},
                &Player{score: 16, id: 1},
                &Player{score: 15, id: 2},
                &Player{score: 14, id: 3},
                &Player{score: 8, id: 4},
                &Player{score: 7, id: 5},
                &Player{score: 3, id: 6},
                &Player{score: 1, id: 7}}

        fmt.Println(BestBalFun(ary))
}
