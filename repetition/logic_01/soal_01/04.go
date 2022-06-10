package main

import "fmt"

func main() {
	fmt.Println("1:")
	satu4(10)
	fmt.Println("\n")

	fmt.Println("2:")
	dua4(10)
	fmt.Println("\n")

	fmt.Println("3:")
	tiga4(10)
	fmt.Println("\n")

	fmt.Println("4:")
	empat4(10)
	fmt.Println("\n")

	fmt.Println("5:")
	lima4(10)
	fmt.Println("\n")

	fmt.Println("6:")
	enam4(10)
	fmt.Println("\n")

	fmt.Println("7:")
	tujuh4(10)
	fmt.Println("\n")

	fmt.Println("8:")
	delapan4(10)
	fmt.Println("\n")

	fmt.Println("9:")
	sembilan4(10)
	fmt.Println("\n")

	fmt.Println("10:")
	sepuluh4(5)
}

func satu4(n int) {
	a := 1

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func dua4(n int) {
	a := 6

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func tiga4(n int) {
	a := 11

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func empat4(n int) {
	a := 16

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func lima4(n int) {
	a := 21

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func enam4(n int) {
	x := 25

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(x, "\t")
			x--
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func tujuh4(n int) {
	a := 20

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(a, "\t")
			a--
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func delapan4(n int) {
	a := 15

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(a, "\t")
			a--
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func sembilan4(n int) {
	a := n

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(a, "\t")
			a--
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func sepuluh4(n int) {
	a := n

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(a, "\t")
			a--
		} else {
			fmt.Print("batas", "\t")
		}
	}
}
