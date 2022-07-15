package main

import (
	"fmt"
)

func main() {
	fmt.Println("1:")
	satu8(12)
	fmt.Println("\n")

	fmt.Println("2:")
	dua8(10)
	fmt.Println("\n")

	fmt.Println("3:")
	tiga8(10)
	fmt.Println("\n")

	fmt.Println("4:")
	empat8(10)
	fmt.Println("\n")

	fmt.Println("5:")
	lima8(10)
	fmt.Println("\n")

	fmt.Println("6:")
	enam8(10)
	fmt.Println("\n")

	fmt.Println("7:")
	tujuh8(10)
	fmt.Println("\n")

	fmt.Println("8:")
	delapan8(10)
	fmt.Println("\n")

	fmt.Println("9:")
	sembilan8(10)
	fmt.Println("\n")

	fmt.Println("10:")
	sepuluh8(10)
}

func satu8(n int) {
	a := 1
	b := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a *= 2
			fmt.Print(a, "\t")
		} else {
			b *= 3
			fmt.Print(b, "\t")
		}
	}
}

func dua8(n int) {
	a := 1
	b := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a *= 3
			fmt.Print(a, "\t")
		} else {
			b *= 2
			fmt.Print(b, "\t")
		}
	}
}

func tiga8(n int) {
	a := 1
	b := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a *= 2
			fmt.Print(a, "\t")
		} else {
			b *= 5
			fmt.Print(b, "\t")
		}
	}
}

func empat8(n int) {
	a := 1
	b := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a *= 5
			fmt.Print(a, "\t")
		} else {
			b *= 2
			fmt.Print(b, "\t")
		}
	}
}

func lima8(n int) {
	a := 1
	b := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a *= 3
			fmt.Print(a, "\t")
		} else {
			b *= 4
			fmt.Print(b, "\t")
		}
	}
}

func enam8(n int) {
	a := 1
	b := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a *= 4
			fmt.Print(a, "\t")
		} else {
			b *= 3
			fmt.Print(b, "\t")
		}
	}
}

func tujuh8(n int) {
	a := 1
	b := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a *= 2
			fmt.Print(a, "\t")
		} else {
			b *= 4
			fmt.Print(b, "\t")
		}
	}
}

func delapan8(n int) {
	a := 1
	b := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a *= 4
			fmt.Print(a, "\t")
		} else {
			b *= 2
			fmt.Print(b, "\t")
		}
	}
}

func sembilan8(n int) {
	a := 1
	b := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a *= 3
			fmt.Print(a, "\t")
		} else {
			b *= 5
			fmt.Print(b, "\t")
		}
	}
}

func sepuluh8(n int) {
	a := 1
	b := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a *= 5
			fmt.Print(a, "\t")
		} else {
			b *= 3
			fmt.Print(b, "\t")
		}
	}
}
