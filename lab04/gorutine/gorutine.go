package gorutine

import "fmt"
 
func StupidGorutine() {
     
    for i := 1; i < 7; i++{
        // создаем горутину для лямбда функции
        go func(n int){
            result := 1
            for j := 1; j <= n; j++{
                result *= j
            }
            fmt.Println(n, "-", result)
        }(i)
    }
   // fmt.Scanln()
    fmt.Println("The End")
}

func SyncedGorutine() {
    results := make(map[int]int)
    structCh := make(chan struct{})
   
    go factorial(5, structCh, results)
     
    <-structCh        // ожидаем закрытия канала structCh
     
    for i, v := range results{
        fmt.Println(i, " - ", v)
    }
}

func factorial(n int, ch chan struct{}, results map[int]int){
    defer close(ch)     // закрываем канал после завершения горутины
    result := 1
    for i:=1; i <= n; i++{
        result *= i
        results[i] = result
    }
}

func StreamGorutine(){
    fmt.Printf("\n")

    intCh := make(chan int)

    go streamFactorial(7, intCh)

    for num := range intCh { // сила полиморфизмо
        fmt.Println(num)
    }
}

func streamFactorial(n int, ch chan int) {
    defer close(ch)
    result := 1
    for i := 1; i <= n; i++{
        result *= i
        ch <- result    
    }
}