package main

import "fmt"


func main() {
	groupArrays := []int{2,4,5,7,6,3,1,8,9,10}
	newArray := []int{}

	for _, value := range groupArrays {
		cuadradoValue := Cuadrado(value)
		newArray = append(newArray, cuadradoValue)
	}

	fmt.Println(newArray)
}

func Cuadrado(number int) int {
	var r, j int
	j = number

	for i := j; i > 0; i-- {
		r += number
	}

	return r
}