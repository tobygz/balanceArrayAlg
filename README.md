# balanceArrayAlg<br>
the best alg to split an array to double<br>
<br>
go run main.go<br>
<br>
src array: [18 16 15 14 8 7 3 1]<br>
normal process<br>
---------------------<br>
[18 14 7 3] 42<br>
[16 15 8 1] 40<br>

start improve process<br>
---------------------<br>
ary: [16 14 7 3]  ary1: [18 15 8 1]  swNum: 1  diff: 2<br>
ary: [18 15 7 3]  ary1: [16 14 8 1]  swNum: 1  diff: 4<br>
ary: [18 14 8 3]  ary1: [16 15 7 1]  swNum: 1  diff: 4<br>
ary: [18 14 7 1]  ary1: [16 15 8 3]  swNum: 1  diff: 2<br>
bestAry: [16 15 7 3] , bestAry1: [18 14 8 1]  diff: 0<br>
ary: [16 15 7 3]  ary1: [18 14 8 1]  swNum: 2  diff: 0<br>
finnal result: [16 15 7 3] ,  [18 14 8 1]  diff: 0<br>
