package main

import "fmt"

func main() {
   var a , b , sum int
   fmt.Scan(&a)
    for i := 1; a < 5; i++ {
        fmt.Scan(&b) 
        if 10<=b && b%8==0 && b<=99 
            sum+=b
            fmt.Println(sum)
    }
}