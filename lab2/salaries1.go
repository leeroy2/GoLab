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


type avSalary struct {
    Average int `json:"average_salary"`
}

type bigSal struct {
    Biggest string `json: "biggest_salary"`
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
var empPerTitle map[string]int
var people []Person
var goldenBoys []Person
var maxSal int

func main() {
    csvFile, _ := os.Open("salaries.csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))
    var p Person
    //var q Person
    var v avSalary
    var salaries []int
    empPerTitle = make(map[string]int)
    num1 := 0
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        num1++
        if num1 == 1 {
            continue
        }
        num1++
        q.Salary, _ = strconv.Atoi(line[3])
        salaries = append(salaries, q.Salary)
    }
    maxSal = Calculate(salaries, "max")

    //empPerTitle = make(map[string]int)
    num := 0
    for {
        line
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
        salaries = append(salaries, p.Salary)
        // Fill values of map
        _, ok := empPerTitle[p.Title]
        if ok == true {
                empPerTitle[p.Title]++
        } else {
                empPerTitle[p.Title] = 1
        }

    }
    v.Average = Calculate(salaries, "average")
    v.avSalary2JSON()
    printJson(empPerTitle)
    maxSal = Calculate(salaries, "max")
    fmt.Println(maxSal)

    peopleJson, _ := json.Marshal(people)
    fmt.Println(string(peopleJson))
    //goldenJson, _ := json.Marshal(goldenBoys)
}
