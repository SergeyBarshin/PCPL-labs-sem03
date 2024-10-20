package mutex

import (
	"fmt"
	"sync"
	"time"
)


var counter int = 0; // общий ресурс

func ExampleWithoutMutex () {
	ch := make(chan bool)

	for i := 1; i < 5; i++ {
		go workWithoutMutex(i, ch)
	}

	// ожидаем завершения всех горутин
	for i := 1; i < 5; i++ {
		<-ch
	}
	fmt.Println("The End")
	// в таком случае вывод будет не 1->5, для каждой горутины. Значения могут путаться
}

func workWithoutMutex(num int, ch chan bool) {
	counter = 0
	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Gorutine", num, "-" ,counter)
	}
	ch <- true
}

func ExampleMutex () {
    ch := make(chan bool)       // канал
    var mutex sync.Mutex        // определяем мьютекс
    for i := 1; i < 5; i++{
        go workMutex(i, ch, &mutex)
    }
     
    for i := 1; i < 5; i++{
        <-ch
    }
     
    fmt.Println("The End")
}

func workMutex (number int, ch chan bool, mutex *sync.Mutex){
    mutex.Lock()    // блокируем доступ к переменной counter
    counter = 0
    for k := 1; k <= 5; k++{
        counter++
        fmt.Println("Goroutine", number, "-", counter)
    }
    mutex.Unlock()  // деблокируем доступ
    ch <- true
}

func WaitGroupe() {
	var wg sync.WaitGroup
    wg.Add(2)       // в группе две горутины
    work := func(id int) {
        defer wg.Done()
        fmt.Printf("Горутина %d начала выполнение \n", id)
        time.Sleep(2 * time.Second)
        fmt.Printf("Горутина %d завершила выполнение \n", id)
   }

   // вызываем горутины
   go work(1)
   go work(2)

   wg.Wait()        // ожидаем завершения обоих горутин
   fmt.Println("Горутины завершили выполнение")
}