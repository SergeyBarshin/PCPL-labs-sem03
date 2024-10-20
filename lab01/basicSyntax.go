package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	// создание переменных, вывод
	var a float64 = 9;
	b := 2;
	var str string
	str = "это строка, а после код символа X:"
	

	fmt.Println(float64(a / float64(b)))
	fmt.Println(str, 'X')

	fmt.Print('1'/2, "\n")
	fmt.Printf("%c\n",'1'+4)

	fmt.Println(miltiple(12, 12))


	// аргументы командной строки
	s2, sep2 := "", ""
	for _, arg := range(os.Args[1:]) {
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(s2)


	// ввод указетели
	fmt.Println("Введите строку и я ее на разверну:")
	var str2 string
	fmt.Scan(&str2)
	reverseStrigns(&str2)
	fmt.Println(str2)

	/*
		массив - обычный массив
		срезы - динамические массивы
		карта - хеш мап
	*/
	// МАССИВЫ
	var arr[5] int
	arr2 := [5]int{1,2,3,4,5}
	fmt.Println(arr, arr2)

	// СРЕЗЫ
	slice := make([]int, 5)  // Создаёт срез из 5 элементов, инициализированных нулями
	slice = append(slice, 6, 7)
	fmt.Println(slice)

	dst := slice[5:7]
	fmt.Println(dst)
	copy(dst, slice)
	fmt.Println(dst)

	dst2 := append(dst, 12, 23)
	fmt.Println(dst2, reflect.TypeOf(slice), reflect.TypeOf(arr), reflect.TypeOf(dst))

	// МАПЫ
	m := map[string]int {}
	fmt.Println(m)
	m["apple"] = 2
	m["banana"] = 5
	m["apple2"] = 4

	fmt.Println(m, m["apple"])

	for key, value := range m {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }

	if _, ok:= m["nonExist"]; !ok {
		fmt.Println("Такого ключа нет")
	}

}

// функции
func miltiple(a int,b int) int {
	res := a
	for i := 1; i < b; i++ {
		res += a
	}
	return res
}

func reverseStrigns(s* string){
	runes := []rune(*s) // стрез рун(чаров)

	for i,j := 0, len(runes)-1; i < j; i,j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*s = string(runes)
}