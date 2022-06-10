package main

import "fmt"

func main() {
	fmt.Println("1 :")
	satu(10)
	fmt.Println("\n")

	fmt.Println("2 :")
	dua(10)
	fmt.Println("\n")

	fmt.Println("3:")
	tiga(10)
	fmt.Println("\n")

	fmt.Println("4:")
	empat(10)
	fmt.Println("\n")

	fmt.Println("5:")
	lima(10)
	fmt.Println("\n")

	fmt.Println("6:")
	enam(10)
	fmt.Println("\n")

	fmt.Println("7:")
	tujuh(10)
	fmt.Println("\n")

	fmt.Println("8:")
	delapan(10)
	fmt.Println("\n")

	fmt.Println("9:")
	sembilan(8)
	fmt.Println("\n")

	fmt.Println("10:")
	sepuluh(6)
}

func satu(n int) {
	a := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func dua(n int) {
	a := 6

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func tiga(n int) {
	a := 11

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func empat(n int) {
	a := 16

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func lima(n int) {
	a := 21

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func enam(n int) {
	x := 25

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(x, "\t")
			x--
		} else {
			fmt.Print("Batas", "\t")
		}
	}
}

func tujuh(n int) {
	a := 20

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a--
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func delapan(n int) {
	a := 15

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a--
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func sembilan(n int) {
	a := n

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a--
		} else {
			fmt.Print("batas", "\t")
		}
	}
}

func sepuluh(n int) {
	a := n

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a--
		} else {
			fmt.Print("batas", "\t")
		}
	}
}
