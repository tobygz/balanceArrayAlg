# balanceArrayAlg
the best alg to split an array to double

go run main.go

src array: [18 16 15 14 8 7 3 1]
normal process
---------------------
[18 14 7 3] 42
[16 15 8 1] 40

start improve process
---------------------
ary: [16 14 7 3]  ary1: [18 15 8 1]  swNum: 1  diff: 2
ary: [18 15 7 3]  ary1: [16 14 8 1]  swNum: 1  diff: 4
ary: [18 14 8 3]  ary1: [16 15 7 1]  swNum: 1  diff: 4
ary: [18 14 7 1]  ary1: [16 15 8 3]  swNum: 1  diff: 2
bestAry: [16 15 7 3] , bestAry1: [18 14 8 1]  diff: 0
ary: [16 15 7 3]  ary1: [18 14 8 1]  swNum: 2  diff: 0
finnal result: [16 15 7 3] ,  [18 14 8 1]  diff: 0
