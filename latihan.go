package main

import (
	"fmt"
	"math"
)
func amstrong(masukan int, pangkat *int){
    var i, j, k, hasil, final int
    j = masukan
    *pangkat = masukan
    for *pangkat > 0{
        *pangkat = *pangkat/10
        i++
    }
    for masukan > 0{
        k = masukan%10
        hasil = int(math.Pow(float64(k), float64(i)))
        final = final + hasil
        masukan = masukan/10
    }
    if j == final{
        fmt.Print(true)
    }else{
        fmt.Print(false)
    }
}

func main(){
    var x, y int
    fmt.Scan(&x)
    amstrong(x, &y)
}