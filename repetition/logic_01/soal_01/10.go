package main

import (
	"fmt"
)

func main() {
	fmt.Println("1:")
	satu10(12)
	fmt.Println("\n")

	fmt.Println("2:")
	dua10(10)
	fmt.Println("\n")

	fmt.Println("3:")
	tiga10(10)
	fmt.Println("\n")

	fmt.Println("4:")
	empat10(10)
	fmt.Println("\n")

	fmt.Println("5:")
	lima10(10)
	fmt.Println("\n")

	fmt.Println("6:")
	enam10(10)
	fmt.Println("\n")

	fmt.Println("7:")
	tujuh10(10)
	fmt.Println("\n")

	fmt.Println("8:")
	delapan10(10)
	fmt.Println("\n")

	fmt.Println("9:")
	sembilan10(10)
	fmt.Println("\n")

	fmt.Println("10:")
	sepuluh10(10)
}

func satu10(n int) {
	a := 1
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 1
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 2
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 3
		} else {
			fmt.Print(999, "\t")
		}
	}
}

func dua10(n int) {
	a := 1
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 2
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 3
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 1
		} else {
			fmt.Print(999, "\t")
		}
	}
}

func tiga10(n int) {
	a := 1
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 3
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 1
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 2
		} else {
			fmt.Print(999, "\t")
		}
	}
}

func empat10(n int) {
	a := 2
	b := 3
	c := 4

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 3
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 2
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 1
		} else {
			fmt.Print(999, "\t")
		}
	}
}

func lima10(n int) {
	a := 2
	b := 3
	c := 4

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 1
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 2
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 3
		} else {
			fmt.Print(999, "\t")
		}
	}
}

func enam10(n int) {
	a := 2
	b := 3
	c := 4

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 3
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 1
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 2
		} else {
			fmt.Print(999, "\t")
		}
	}
}

func tujuh10(n int) {
	a := 2
	b := 3
	c := 4

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 3
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 3
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 3
		} else {
			fmt.Print(999, "\t")
		}
	}
}

func delapan10(n int) {
	a := 2
	b := 3
	c := 4

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 2
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 2
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 2
		} else {
			fmt.Print(999, "\t")
		}
	}
}

func sembilan10(n int) {
	a := 2
	b := 3
	c := 4

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 3
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 1
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 5
		} else {
			fmt.Print(999, "\t")
		}
	}
}

func sepuluh10(n int) {
	a := 1
	b := 3
	c := 5

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 3
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 2
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 1
		} else {
			fmt.Print(999, "\t")
		}
	}
}
