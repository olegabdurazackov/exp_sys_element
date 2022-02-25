package main

import(
    "fmt"
    "math/rand"
    "time"
    "github.com/goml/gobrain"
    e "./elem"
)


func main(){
    seed:=time.Now().Unix()
    rand.Seed(seed)
    ff:=&gobrain.FeedForward{}
    ff.Init(216,110,55)
    ff.Train(e.Patterns,1000,0.6,0.4,false)

    in := []float64{
     /*0*/   0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,

      /*50*/  0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,


      /*100*/ 0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,

      /*150*/ 0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0,0,0,0,0,


      /*200*/ 0,0,0,0,0, 0,0,0,0,0,
              0,0,0,0,0, 0}

		len_in:=len(in)
		fmt.Println("Cледует ввести 1 - да или 0 -нет")
		for i:=0;i<len_in;i++{
				fmt.Println(i,e.In[i])
				fmt.Scan(&in[i])
				fmt.Println()
		}
		fmt.Println("Конец опроса. Вычисление..")
		out := ff.Update(in)
    fmt.Println(in)
    fmt.Println(out)
    li:=len(out)
    for i:=0; i<li;i++ {
      if out[i]>0.1{
          fmt.Println("\n",out[i],e.Out[i])
      }
    }
/*
    in = []float64{0, 0, 0, 0}
    out = ff.Update(in)
    fmt.Println(in)
    fmt.Println(out)

    var a,b,c,d float64 
    for i:=0;true;i++{
        fmt.Println("Ввод a,b,c,d")
        fmt.Scan(&a,&b,&c,&d)
        in = []float64{a, b, c, d}
        out = ff.Update(in)
        fmt.Println(in)
        fmt.Println(out)
    }
*/
}
