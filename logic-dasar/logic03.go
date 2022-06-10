package main

import "fmt"

func main() {
	fmt.Println("soal 01")
	logic03soal01(9)

	fmt.Println("soal 02")
	logic03soal02(9)

	fmt.Println("soal 03")
	logic03soal03(9)

	fmt.Println("soal 04")
	logic03soal04(9)

	fmt.Println("soal 05")
	logic03soal05(9)

	Ngide(9)
	Ngide2(9)
}

func logic03soal01(n int) {
	a := 3
	nt := n / 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j >= n-1 && i <= j {
				fmt.Print(a, "\t")
			} else if i+j <= n-1 && i >= j {
				fmt.Print(a, "\t")
			} else {
				fmt.Print("\t")
			}
		}
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
		fmt.Println("\n")
	}
}

func logic03soal02(n int) {
	a := 3
	nt := n / 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j <= n-1 && i <= j {
				fmt.Print(a, "\t")
			} else if i+j >= n-1 && i >= j {
				fmt.Print(a, "\t")
			} else {
				fmt.Print("\t")
			}
		}
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
		fmt.Println("\n")
	}
}

func logic03soal03(n int) {
	a := 3
	nt := n / 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > i && j <= nt {
				fmt.Print(a, "\t")
			} else if i+j >= n && i < nt {
				fmt.Print(a, "\t")
			} else if (i > j && i+j < n-1) && i > nt {
				fmt.Print(a, "\t")
			} else if (j < i && i+j >= n) && j >= nt {
				fmt.Print(a, "\t")
			} else if i == nt {
				fmt.Print(a, "\t")
			} else {
				fmt.Print("\t")
			}
		}
		fmt.Println("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}

func logic03soal04(n int) {
	a := 3
	nt := n / 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				fmt.Print(a, "\t")
			} else if i%4 == 1 && j == n-1 {
				fmt.Print(a, "\t")
			} else if i%4 == 3 && j == 0 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print("\t")
			}
		}
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
		fmt.Println("\n")
	}
}

func logic03soal05(n int) {
	nt := n / 2

	for i := 0; i < n; i++ {
		a := 3
		for j := 0; j < n; j++ {
			if j%2 == 0 {
				fmt.Print(a, "\t")
			} else if j%4 == 1 && i == n-1 {
				fmt.Print(a, "\t")
			} else if j%4 == 3 && i == 0 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print("\t")
			}
			if j < nt {
				a += 3
			} else {
				a -= 3
			}
		}
		fmt.Println("\n")
	}
}

func Ngide(n int) {
	a := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%4 == 0 {
				fmt.Print(a, "\t")
			} else if j%4 == 0 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print("\t")
			}
		}
		a++
		fmt.Print("\n")
	}
}

func Ngide2(n int) {
	a := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i+1 == j {
				fmt.Print(a, "\t")
			} else {
				fmt.Print("\t")
			}
		}
		a++
		fmt.Println("\n")
	}
}
