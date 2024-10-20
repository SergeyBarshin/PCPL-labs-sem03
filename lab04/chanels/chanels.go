package chanels

import "fmt"

func Square(ch chan int){
      
    num := <-ch                 // получаем из канала число
    fmt.Println("num := ", num)
    ch <- num * num             // обратно отправляем квадрат числа
}