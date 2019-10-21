package main

import ("fmt")

func add(a,b float64) float64 {
	return a+b
}
func mutiply(a,b float64) float64 {
	return a*b
}



func main() {
	a,b :=6.6, 5.3
	fmt.Println(add(a,b))
}