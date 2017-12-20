package main

import (
	"fmt"
	"encoding/json"
        "encoding/csv"
	"bufio"
	"io"
	"os"
	"log"
	"strconv"
)

type Person struct {
    Firstname string   `json:"firstname"`
    Surname string  `json:"surname"`
    Title string `json:"title"`
    Salary int `json:salary"`
}

type PersonSlice []Person

func (q PersonSlice) getGoldenBoys() PersonSlice {
	var entries PersonSlice
	for i, v := range q {
		if v.Salary == maxSal {
			entries = append(entries, q[i])
		}
	}
	return entries
}

type avSalary struct {
    Average int `json:"average_salary"`
}

type bigSal struct {
    Biggest_salary PersonSlice `json: "biggest_salary"`
}

func (z bigSal) goldenBoys2JSON() {
    golden2JSON, _ :=json.Marshal(z)
    fmt.Println(string(golden2JSON))
}

func (v avSalary) avSalary2JSON() {
    avSal2JSON, _ := json.Marshal(v)
    fmt.Println(string(avSal2JSON))
}


func Calculate (salary []int, flag string) int {
    var sum,av,max int
    for _, sal := range salary {
        sum += sal
        if sal > max {
            max = sal
        }
    }
    if len(salary) != 0 {
        av = sum/len(salary)
    }
    switch flag {
    case "average":
        return av
    case "max":
        return max
    }
    return 0
}

func printJson(aMap map[string]int) {
        mapB, _ := json.Marshal(aMap)
        fmt.Println(string(mapB))
}

//Global variables
var employeesPerTitle map[string]int
var people []Person
var maxSal int

func main() {
    csvFile, _ := os.Open("salaries.csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))
    var p Person
    var q PersonSlice
    var v avSalary
    var b bigSal
    var salaries []int
    employeesPerTitle = make(map[string]int)
    num := 0
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
	num++
	if num == 1 {
            continue
	}
	num++
	p.Firstname = line[0]
	p.Surname = line[1]
	p.Title = line[2]
	p.Salary, _ = strconv.Atoi(line[3])
	people = append(people, p)
        q = append(q, p)
        salaries = append(salaries, p.Salary)
        // Fill values of map
        _, ok := employeesPerTitle[p.Title]
        if ok == true {
                employeesPerTitle[p.Title]++
        } else {
                employeesPerTitle[p.Title] = 1
        }
    }
    maxSal = Calculate(salaries, "max")
    v.Average = Calculate(salaries, "average")
    v.avSalary2JSON()
    printJson(employeesPerTitle)
    b.Biggest_salary = q.getGoldenBoys()
    b.goldenBoys2JSON() 
}
