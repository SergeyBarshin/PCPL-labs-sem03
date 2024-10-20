/*
Ram - хранилище, работающие через интерфейс
*/

package db

import "fmt"


type IEmployee interface { // общий интерфейс, для полиморфизми структур
	Ping() string
}

type Employee struct {
    Name string
    Age  int
}

type SoftwareEngenere struct {
	Employee
	Language string
}

type DataScientist struct {
	Employee
	IsNlp bool
}

type DevOps struct {
	Employee
	Stack []string
}


func (a SoftwareEngenere) Ping() string {
	return fmt.Sprintf("I am SWE, i code on %s and I my Name is %s", a.Language, a.Employee.Name)
}

func (a DataScientist) Ping() string {
	return fmt.Sprintf("I am DS, is nlp %t and I my Name is %s", a.IsNlp, a.Employee.Name)
}
func (a DevOps) Ping() string {
	return fmt.Sprintf("I am DevOps,I my Name is %s, my stack is %s", a.Employee.Name, a.Stack)
}

func Ping(s IEmployee) {
    fmt.Println(s.Ping())
}

type RamStorage struct {
	data map[string]IEmployee
}

func (db *RamStorage) Create(name string, obj IEmployee) {
    if db.data == nil {
        db.data = make(map[string]IEmployee)
    }

	if _, exists := db.data[name]; exists {
        return
    } else {
        db.data[name] = obj
    }
}

func (db *RamStorage) Update(name string, obj IEmployee) {
    if db.data == nil {
        db.data = make(map[string]IEmployee)
    }
    db.data[name] = obj
}


func (db *RamStorage) Read(name string) IEmployee {
	if _, ok := db.data[name]; !ok{
		return nil
	}		
	return db.data[name]
}

func (db *RamStorage) Delete(name string) bool {
	if _, exists := db.data[name]; exists {
		delete(db.data, name)
        return true
    } else {
        return false
    }
}