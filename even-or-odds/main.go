package main

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenOrOdd(slice)
}

func evenOrOdd(s []int) {
	for _, num := range s {
		if num%2 == 0 {
			println(num, "Even")
		} else {
			println(num, "Odd")
		}
	}
}
