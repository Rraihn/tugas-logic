package main

import "fmt"

func main() {
	fmt.Println("1:")
	satu5(18)
	fmt.Println("\n")

	fmt.Println("2:")
	dua5(12)
	fmt.Println("\n")

	fmt.Println("3:")
	tiga5(12)
	fmt.Println("\n")

	fmt.Println("4:")
	empat5(12)
	fmt.Println("\n")

	fmt.Println("5:")
	lima5(10)
	fmt.Println("\n")

	fmt.Println("6:")
	enam5(10)
	fmt.Println("\n")

	fmt.Println("7:")
	tujuh5(10)
	fmt.Println("\n")

	fmt.Println("8:")
	delapan5(10)
	fmt.Println("\n")

	fmt.Println("9:")
	sembilan5(10)
	fmt.Println("\n")

	fmt.Println("10:")
	sepuluh5(10)
}

func satu5(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("fizz, buzz", "\t")
		} else if i%3 == 0 {
			fmt.Print("fizz", "\t")
		} else if i%5 == 0 {
			fmt.Print("buzz", "\t")
		} else {
			fmt.Print(i, "\t")
		}
	}
}

func dua5(n int) {
	for i := 1; i <= n; i++ {
		if i%4 == 0 && i%3 == 0 {
			fmt.Print("tiga, empat", "\t")
		} else if i%3 == 0 {
			fmt.Print("tiga", "\t")
		} else if i%4 == 0 {
			fmt.Print("empat", "\t")
		} else {
			fmt.Print(i, "\t")
		}
	}
}

func tiga5(n int) {
	a := 1

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%4 == 0 {
			fmt.Print("tiga, empat", "\t")
		} else if i%3 == 0 {
			fmt.Print("tiga", "\t")
		} else if i%4 == 0 {
			fmt.Print("empat", "\t")
		} else {
			fmt.Print(a, "\t")
			a += 2
		}
	}
}

func empat5(n int) {
	a := 1

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%4 == 0 {
			fmt.Print("tiga, empat", "\t")
		} else if i%3 == 0 {
			fmt.Print("tiga", "\t")
		} else if i%4 == 0 {
			fmt.Print("empat", "\t")
		} else {
			fmt.Print(a, "\t")
			a *= 2
		}
	}
}

func lima5(n int) {
	a := 1

	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%5 == 0 {
			fmt.Print("dua, lima", "\t")
		} else if i%2 == 0 {
			fmt.Print("dua", "\t")
		} else if i%5 == 0 {
			fmt.Print("lima", "\t")
		} else {
			fmt.Print(a, "\t")
			a += 2
		}
	}
}

func enam5(n int) {
	a := 1

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("tiga, lima", "\t")
		} else if i%5 == 0 {
			fmt.Print("tiga", "\t")
		} else if i%5 == 0 {
			fmt.Print("lima", "\t")
		} else {
			fmt.Print(a, "\t")
			a += 2
		}
	}
}

func tujuh5(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Print("dua", "\t")
		} else {
			fmt.Print(i, "\t")
		}
	}
}

func delapan5(n int) {
	a := 1

	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			fmt.Print("repeat", "\t")
		} else {
			fmt.Print(a, "\t")
			a += 3
		}
	}
}

func sembilan5(n int) {
	a := 1

	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Print("repeat", "\t")
		} else {
			fmt.Print(a, "\t")
			a += 5
		}
	}
}

func sepuluh5(n int) {
	a := 1

	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%5 == 0 {
			fmt.Print("dua, lima", "\t")
		} else if i%2 == 0 {
			fmt.Print("dua", "\t")
		} else if i%5 == 0 {
			fmt.Print("lima", "\t")
		} else {
			fmt.Print(a, "\t")
			a += 4
		}
	}
}
