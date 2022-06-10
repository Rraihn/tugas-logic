package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("1:")
	satu6(15)
	fmt.Println("\n")

	fmt.Println("2:")
	dua6(10)
	fmt.Println("\n")

	fmt.Println("3:")
	tiga6(10)
	fmt.Println("\n")

	fmt.Println("4:")
	empat6(12)
	fmt.Println("\n")

	fmt.Println("5:")
	lima6(10)
	fmt.Println("\n")

	fmt.Println("6:")
	enam6(12)
	fmt.Println("\n")

	fmt.Println("7:")
	tujuh6(15)
	fmt.Println("\n")

	fmt.Println("8:")
	delapan6(10)
	fmt.Println("\n")

	fmt.Println("9:")
	sembilan6(10)
	fmt.Println("\n")

	fmt.Println("10:")
	sepuluh6(10)
}

func satu6(n int) {
	a := 3
	nt := n / 2

	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			x := math.Pow(float64(i), 2)
			fmt.Print(x, "\t")
		} else {
			fmt.Print(a, "\t")
		}
		if i <= nt {
			a += 3
		} else {
			a -= 3
		}
	}
}

func dua6(n int) {
	a := 3
	nt := n / 2

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			x := math.Pow(float64(i), 2)
			fmt.Print(x, "\t")
		} else {
			fmt.Print(a, "\t")
		}
		if i <= nt {
			a += 2
		} else {
			a -= 2
		}
	}
}

func tiga6(n int) {
	a := 3
	nt := n / 2

	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			x := math.Pow(float64(i), 2)
			fmt.Print(x, "\t")
		} else {
			fmt.Print(a, "\t")
		}
		if i <= nt {
			a += 2
		} else {
			a -= 2
		}
	}
}

func empat6(n int) {
	a := 5
	nt := n / 2

	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			x := math.Pow(float64(i), 2)
			fmt.Print(x, "\t")
		} else {
			fmt.Print(a, "\t")
		}
		if i <= nt {
			a += 3
		} else {
			a -= 3
		}
	}
}

func lima6(n int) {
	a := 3
	nt := n / 2

	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			x := math.Pow(float64(i), 2)
			fmt.Print(x, "\t")
		} else {
			fmt.Print(a, "\t")
		}
		if i <= nt {
			a *= 2
		} else {
			a /= 2
		}
	}
}

func enam6(n int) {
	a := 5
	nt := n / 2

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			x := math.Pow(float64(i), 2)
			fmt.Print(x, "\t")
		} else {
			fmt.Print(a, "\t")
		}
		if i <= nt {
			a += 5
		} else {
			a -= 5
		}
	}
}

func tujuh6(n int) {
	a := 3
	nt := n / 2

	for i := 1; i <= n; i++ {
		if i%5 == 0 {
			x := math.Pow(float64(i), 2)
			fmt.Print(x, "\t")
		} else {
			fmt.Print(a, "\t")
		}
		if i <= nt {
			a += 2
		} else {
			a -= 2
		}
	}
}

func delapan6(n int) {
	a := 1
	nt := n / 2

	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			x := math.Pow(float64(i), 2)
			fmt.Print(x, "\t")
		} else {
			fmt.Print(a, "\t")
		}
		if i <= nt {
			a += 3
		} else {
			a -= 3
		}
	}
}

func sembilan6(n int) {
	a := 5
	nt := n / 2

	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			x := math.Pow(float64(i), 2)
			fmt.Print(x, "\t")
		} else {
			fmt.Print(a, "\t")
		}
		if i <= nt {
			a += 2
		} else {
			a -= 2
		}
	}
}

func sepuluh6(n int) {
	a := 2
	nt := n / 2

	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			x := math.Pow(float64(i), 2)
			fmt.Print(x, "\t")
		} else {
			fmt.Print(a, "\t")
		}
		if i <= nt {
			a *= 2
		} else {
			a /= 2
		}
	}
}
