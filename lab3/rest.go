package main

import (
	"fmt"
	"encoding/json"
        "encoding/csv"
	"bufio"
	"io"
        "io/ioutil"
	"os"
	"log"
	"strconv"
        "net/http"
        "net"
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

func (z bigSal) goldenBoys2JSON() string {
    golden2JSON, _ :=json.Marshal(z)
    return string(golden2JSON)
}

func (v avSalary) avSalary2JSON() string {
    avSal2JSON, _ := json.Marshal(v)
    return string(avSal2JSON)
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

func printJson(aMap map[string]int) string {
        mapB, _ := json.Marshal(aMap)
        //fmt.Println(string(mapB))
        return string(mapB)
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
    var user Person
    var b bigSal
    var salaries []int
    found := 0
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
    average := v.avSalary2JSON()
    empPerTitle := printJson(employeesPerTitle)
    b.Biggest_salary = q.getGoldenBoys()
    goldenBoys := b.goldenBoys2JSON()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path[1:]
        ip,_,_ := net.SplitHostPort(r.RemoteAddr)
        switch path {
        case "average":
            log.Println("There is a new call at " + r.URL.String() + " from" + ip)
            fmt.Fprintf(w, average)
        case "employees":
            log.Println("There is a new call at " + r.URL.String() + " from" + ip)
            fmt.Fprintf(w, empPerTitle)
        case "big":
            log.Println("There is a new call at " + r.URL.String() + " from" + ip)
            fmt.Fprintf(w, goldenBoys)
        case "employee":
            log.Println("There is a new call at " + r.URL.String() + " from" + ip)
            body, err := ioutil.ReadAll(r.Body)
            if err != nil {
	        http.Error(w, err.Error(), 500) 
	    }

            if len(body) > 0 {
                json.Unmarshal(body, &user)
                for _, value := range people {
                    if value.Surname == user.Surname {
                        found = 1
                        response, _ :=json.Marshal(value)
                        fmt.Fprintf(w, "%s", response)
                        log.Println("Searching for " + user.Surname)
                        log.Println(user.Surname + " exists in our file(s) from" + ip)
                    }
                }
                if found == 0 {
                    http.Error(w, "Sorry! The requested value does not exists", 404)
                    log.Println(user.Surname + " does not exists in our file(s) from" + ip)
                }
            } else {
                fmt.Fprintf(w, "Oops! Request body is empty!")
                log.Println("Request body is empty at " + r.URL.String() )
            }
        default:
            http.Error(w, "Error 404! page not found", 404)
        }
    })

    http.ListenAndServe(":8084", nil)
}
