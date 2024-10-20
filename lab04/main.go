package main

import (
	"concurs/chanels"
	"concurs/gorutine"
	"concurs/mutex"
	"fmt"
)

func main() {
	//gorutine.StupidGorutine() // недерменированная функция.перемешку

	// Каналы
	intCh := make(chan int)
	go chanels.Square(intCh)
	intCh <- 4 // без этой строки будет deadlock, так как и main b Square будут ждать значение в канал
	fmt.Println("result := ", <- intCh)
	fmt.Printf("\n")

	// синхронизированн
	gorutine.SyncedGorutine()

	// Потоковая передача
	gorutine.StreamGorutine()

	fmt.Printf("\n")
	mutex.ExampleWithoutMutex()
	fmt.Printf("\n")
	mutex.ExampleMutex()
	fmt.Printf("\n")

	mutex.WaitGroupe()
}