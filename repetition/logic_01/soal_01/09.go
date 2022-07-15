package main

import (
	"fmt"
)

func main() {
	fmt.Println("1:")
	satu9(12)
	fmt.Println("\n")

	fmt.Println("2:")
	dua9(10)
	fmt.Println("\n")

	fmt.Println("3:")
	tiga9(10)
	fmt.Println("\n")

	fmt.Println("4:")
	empat9(10)
	fmt.Println("\n")

	fmt.Println("5:")
	lima9(10)
	fmt.Println("\n")

	fmt.Println("6:")
	enam9(10)
	fmt.Println("\n")

	fmt.Println("7:")
	tujuh9(10)
	fmt.Println("\n")

	fmt.Println("8:")
	delapan9(10)
	fmt.Println("\n")

	fmt.Println("9:")
	sembilan9(10)
	fmt.Println("\n")

	fmt.Println("10:")
	sepuluh9(10)
}

func satu9(n int) {
	a := 1
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 2
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 3
		} else {
			fmt.Print(a, "\t")
		}
	}
}

func dua9(n int) {
	a := 1
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 3
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 2
		} else {
			fmt.Print(a, "\t")
		}
	}
}

func tiga9(n int) {
	a := 3
	b := 2
	c := 1

	for i := 1; i <= n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 2
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 3
		} else {
			fmt.Print(a, "\t")
		}
	}
}

func empat9(n int) {
	a := 3
	b := 2
	c := 1

	for i := 1; i <= n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 3
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 2
		} else {
			fmt.Print(a, "\t")
		}
	}
}

func lima9(n int) {
	a := 2
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 2
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 3
		} else {
			fmt.Print(a, "\t")
		}
	}
}

func enam9(n int) {
	a := 2
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 3
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 2
		} else {
			fmt.Print(a, "\t")
		}
	}
}

func tujuh9(n int) {
	a := 3
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 2
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 3
		} else {
			fmt.Print(a, "\t")
		}
	}
}

func delapan9(n int) {
	a := 3
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 3
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 2
		} else {
			fmt.Print(a, "\t")
		}
	}
}

func sembilan9(n int) {
	a := 1
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 3
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 4
		} else {
			fmt.Print(a, "\t")
		}
	}
}

func sepuluh9(n int) {
	a := 1
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 4
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 3
		} else {
			fmt.Print(a, "\t")
		}
	}
}
