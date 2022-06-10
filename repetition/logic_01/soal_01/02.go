package main

import "fmt"

func main() {
	fmt.Println("1:")
	Asatu(10)
	fmt.Println("\n")

	fmt.Println("2:")
	Adua(10)
	fmt.Println("\n")

	fmt.Println("3:")
	Atiga(10)
	fmt.Println("\n")

	fmt.Println("4:")
	Aempat(10)
	fmt.Println("\n")

	fmt.Println("5:")
	Alima(10)
	fmt.Println("\n")

	fmt.Println("6:")
	Aenam(10)
	fmt.Println("\n")

	fmt.Println("7:")
	Atujuh(10)
	fmt.Println("\n")

	fmt.Println("8:")
	Adelapan(10)
	fmt.Println("\n")

	fmt.Println("9:")
	Asembilan(8)
	fmt.Println("\n")

	fmt.Println("10:")
	Asepuluh(6)
}

func Asatu(n int) {
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

func Adua(n int) {
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

func Atiga(n int) {
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

func Aempat(n int) {
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

func Alima(n int) {
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

func Aenam(n int) {
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

func Atujuh(n int) {
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

func Adelapan(n int) {
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

func Asembilan(n int) {
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

func Asepuluh(n int) {
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
