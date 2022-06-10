package logic_03

import (
	"fmt"
	array2 "github.com/Arraf18/logic-dasar/array"
)

func Logic03Soal01(n int) {
	angka := 3
	nt := n / 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= j && i+j <= n-1 {
				fmt.Print(angka, "\t")
			} else if i <= j && i+j >= n-1 {
				fmt.Print(angka, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Println("\n")
		if i < nt {
			angka += 3
		} else {
			angka -= 3
		}
	}
}

func Logic03Soal02(n int) {
	angka := 3
	nt := n / 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j && i+j <= n-1 {
				fmt.Print(angka, "\t")
			} else if i >= j && i+j >= n-1 {
				fmt.Print(angka, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Println("\n")
		if i < nt {
			angka += 3
		} else {
			angka -= 3
		}
	}
}

func Logic03Soal03(n int) {
	// 1. create array
	array := array2.NewNumberArray(n, n)

	// 2. isi array
	// loop baris
	angka := 3 // type nya int
	for i := 0; i < len(array); i++ {
		nt := n / 2
		// loop kolom
		for j := 0; j < len(array[i]); j++ {
			// isi array
			if i < j && j <= nt {
				array[i][j] = int32(angka) // angka di convert ke int32
			} else if i > j && j >= nt {
				array[i][j] = int32(angka) // angka di convert ke int32
			} else if i+j < n-1 && i >= nt {
				array[i][j] = int32(angka)
			} else if i+j > n-1 && i <= nt {
				array[i][j] = int32(angka)
			}
		}

		// jika baris kurang dari nilai tengah
		if i < nt {
			angka += 3 // angka = angka + 3
		} else {
			angka -= 3
		}
	}

	// 3. print array
	array2.PrintNumberArrayZeroEmpty(array)
}

func Logic03Soal04(n int) {
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
		fmt.Print("\n")
	}
}
