package main

import (
	"fmt"
)

func main() {
	fmt.Println("1:")
	satu7(15)
	fmt.Println("\n")

	fmt.Println("2:")
	dua7(10)
	fmt.Println("\n")

	fmt.Println("3:")
	tiga7(10)
	fmt.Println("\n")

	fmt.Println("4:")
	empat7(10)
	fmt.Println("\n")

	fmt.Println("5:")
	lima7(10)
	fmt.Println("\n")

	fmt.Println("6:")
	enam7(10)
	fmt.Println("\n")

	fmt.Println("7:")
	tujuh7(10)
	fmt.Println("\n")

	fmt.Println("8:")
	delapan7(10)
	fmt.Println("\n")

	fmt.Println("9:")
	sembilan7(10)
	fmt.Println("\n")

	fmt.Println("10:")
	sepuluh7(12)
}

func satu7(n int) {
	m := 4
	num := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(m, "\t")
		} else {
			num += 1
			fmt.Print(num, "\t")
		}
	}
}

func dua7(n int) {
	m := 4
	num := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(m, "\t")
		} else {
			num += 2
			fmt.Print(num, "\t")
		}
	}
}

func tiga7(n int) {
	m := 2
	num := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(m, "\t")
		} else {
			num += 2
			fmt.Print(num, "\t")
		}
	}
}

func empat7(n int) {
	m := 2
	num := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(m, "\t")
		} else {
			num += 3
			fmt.Print(num, "\t")
		}
	}
}

func lima7(n int) {
	m := 3
	num := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(m, "\t")
		} else {
			num += 3
			fmt.Print(num, "\t")
		}
	}
}

func enam7(n int) {
	m := 3
	num := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(m, "\t")
		} else {
			num += 2
			fmt.Print(num, "\t")
		}
	}
}

func tujuh7(n int) {
	m := 1
	num := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(m, "\t")
		} else {
			num += 3
			fmt.Print(num, "\t")
		}
	}
}

func delapan7(n int) {
	m := 1
	num := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(m, "\t")
		} else {
			num += 4
			fmt.Print(num, "\t")
		}
	}
}

func sembilan7(n int) {
	m := 5
	num := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(m, "\t")
		} else {
			num += 2
			fmt.Print(num, "\t")
		}
	}
}

func sepuluh7(n int) {
	m := 5
	num := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(m, "\t")
		} else {
			num += 3
			fmt.Print(num, "\t")
		}
	}
}
