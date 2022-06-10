package main

import "fmt"

func main() {
	fmt.Println("1 :")
	satu3(10)
	fmt.Println("\n")

	fmt.Println("2 :")
	dua3(10)
	fmt.Println("\n")

	fmt.Println("3:")
	tiga3(10)
	fmt.Println("\n")

	fmt.Println("4:")
	empat3(10)
	fmt.Println("\n")

	fmt.Println("5:")
	lima3(10)
	fmt.Println("\n")

	fmt.Println("6:")
	enam3(10)
	fmt.Println("\n")

	fmt.Println("7:")
	tujuh3(10)
	fmt.Println("\n")

	fmt.Println("8:")
	delapan3(10)
	fmt.Println("\n")

	fmt.Println("9:")
	sembilan3(8)
	fmt.Println("\n")

	fmt.Println("10:")
	sepuluh3(5)
}

func satu3(n int) {
	a := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("satu", "\t")
		}
	}
}

func dua3(n int) {
	a := 6

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("satu", "\t")
		}
	}
}

func tiga3(n int) {
	a := 11

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("satu", "\t")
		}
	}
}

func empat3(n int) {
	a := 16

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("satu", "\t")
		}
	}
}

func lima3(n int) {
	a := 21

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print("satu", "\t")
		}
	}
}

func enam3(n int) {
	x := 25

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(x, "\t")
			x--
		} else {
			fmt.Print("satu", "\t")
		}
	}
}

func tujuh3(n int) {
	a := 20

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a--
		} else {
			fmt.Print("satu", "\t")
		}
	}
}

func delapan3(n int) {
	a := 15

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a--
		} else {
			fmt.Print("satu", "\t")
		}
	}
}

func sembilan3(n int) {
	a := n

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a--
		} else {
			fmt.Print("satu", "\t")
		}
	}
}

func sepuluh3(n int) {
	a := n

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a--
		} else {
			fmt.Print("satu", "\t")
		}
	}
}
