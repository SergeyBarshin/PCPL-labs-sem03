/*
в go не структуры, не классы. Не поддерживают наследование, вместо него композиция
Методы: типо функций в классе,

если поле с маленькой буквы, то оно private, иначе public
*/
package main

import (
	"fmt"

	"dbProject/db" // импортируем наш модуль, где "example.com/project" — это модуль из go.mod
)


func main() {
	swe := db.SoftwareEngenere{Employee: db.Employee{Name: "name", Age: 19}, Language: "go-lang"}
	ml := db.DataScientist{Employee: db.Employee{Name: "name2", Age: 20}, IsNlp: true}
	devOps := db.DevOps{Employee: db.Employee{Name: "name3", Age: 21}}
	devOps.Stack = append(devOps.Stack, "k8s", "terraform")

	//swe.Employee.

	db.Ping(swe) // вызываем фуекцию из модуля, в данном случае db не переменная
	db.Ping(ml)
	db.Ping(devOps) // не сработает, если у DevOps нет такого метода
	fmt.Println(devOps.Stack, devOps.Employee.Name)
	swe.Ping()

	fmt.Println("\n\nТестируем напрямую")
	dp1 := db.RamStorage{}
	// не сработает потому что data с маленькой буквы
	// dp1.data = make(map[string]db.IEmployee) // "выделяем память"
	// dp1.data["id1"] = swe
	fmt.Println(dp1)

	dataBase := db.RamStorage{}

	fmt.Println("\n\nТестируем CREATE")
	dataBase.Create("id1", ml)
	dataBase.Create("id2", devOps)
	fmt.Println(devOps)

	fmt.Println("\n\nТестируем READ")
	obj := dataBase.Read("id1")
	fmt.Println(obj)
	if(obj != nil) {fmt.Println(obj.Ping())}
	obj2 := dataBase.Read("id3")
	if(obj2 != nil) {fmt.Println(obj2.Ping())}


	fmt.Println("\n\nТестируем UPDATE")
	fmt.Println(dataBase)
	dataBase.Update("id1", swe)
	fmt.Println(dataBase)

	fmt.Println("\n\nТестируем DELETE")
	fmt.Println(dataBase)
	dataBase.Delete("id2")
	fmt.Printf("%s \n\n", dataBase)

	// еще немного полиморфизма
	employes := []db.IEmployee{swe, devOps, ml}
	for _, el := range employes {
		fmt.Println(el.Ping())
	}
}
