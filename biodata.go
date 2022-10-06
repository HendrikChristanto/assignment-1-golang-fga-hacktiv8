package main
 
import (
	"os"
    "fmt"
	"strconv"
)

type Student struct {
	Name 	string
	Address string
	Job 	string
	Reason	string
}

var students = []Student{
	{Name: "Hendrik", Address: "JL. Majapahit", Job: "Student", Reason: "Want to learn backend"},
	{Name: "Joko", Address: "JL. Ahmad Yani", Job: "Student", Reason: "Want to learn backend"},
	{Name: "Maya", Address: "JL. Basuki Rahmat", Job: "Student", Reason: "Want to learn backend"},
	{Name: "Tingkir", Address: "JL. Kertajaya", Job: "Student", Reason: "Want to learn backend"},
	{Name: "Bagus", Address: "JL. Mulyorejo", Job: "Student", Reason: "Want to learn backend"},
}
 
func main() {
    printStudent(os.Args[1])
}

func printStudent(numberInput string) {
	if number, err := strconv.Atoi(numberInput); err != nil {
        fmt.Println("Please input a number")
    } else if number > 0 && number <= len(students) {
		fmt.Printf("Data : %+v\n", students[number - 1])
	} else {
		fmt.Println("Out of number")
	}
}