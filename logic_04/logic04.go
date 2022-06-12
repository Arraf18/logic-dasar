package logic_04

import (
	array2 "github.com/Arraf18/logic-dasar/array"
)

func Logic04Soal01(n int) {
	array := array2.NewNumberArray(n, n) // untuk membuat array
	a := 1
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if i%4 == 0 { //untuk membuat deret angka dari kiri ke kanan
				array[i][j] = int32(a)
				a++
			} else if i%4 == 1 && j == n-1 { // untuk membuat kotak kosong pada baris 1 dan berisi pada kolol 8
				array[i][j] = int32(a)
				a++
			} else if i%4 == 2 { // untuk membuat baris ke 3 dan 6 bergerak secara mundur
				array[i][n-1-j] = int32(a)
				a++
			} else if i%4 == 3 && j == 0 { // untuk membuat isi pada kolom 0
				array[i][j] = int32(a)
				a++
			}
		}
	}
	array2.PrintNumberArrayZeroEmpty(array) // untuk print array2dimensi
}

func Logic04Soal02(n int) {
	// create array
	array := array2.NewNumberArray(n, n)
	// initial angka dari 1
	angka := 1
	// loop baris
	for b := 0; b < n; b++ {
		// loop kolomg
		for k := 0; k < n; k++ {
			// jika baris 0
			if b%4 == 0 {
				array[k][b] = int32(angka)
				angka += 3
			} else if b%4 == 1 && k == n-1 {
				array[k][b] = int32(angka)
				angka += 3
			} else if b%4 == 2 {
				array[n-1-k][b] = int32(angka)
				angka += 3
			} else if b%4 == 3 && k == 0 {
				array[k][b] = int32(angka)
				angka += 3
			}
		}
	}
	// print array
	array2.PrintNumberArrayZeroEmpty(array)
}

func Logic04Soal03(n int) {
	// create array
	array := array2.NewNumberArray(n, n)
	// initial angka dari 1
	angka := 1
	// loop baris
	for i := 0; i < n; i++ {
		// loop kolomg
		for j := 0; j < n; j++ {
			// jika baris 0
			if i%4 == 0 {
				array[n-1-j][i] = int32(angka)
				angka += 3
			} else if i%4 == 1 && j == 0 {
				array[j][i] = int32(angka)
				angka += 3
			} else if i%4 == 2 {
				array[j][i] = int32(angka)
				angka += 3
			} else if i%4 == 3 && j == n-1 {
				array[j][i] = int32(angka)
				angka += 3
			}
		}
	}
	// print array
	array2.PrintNumberArrayZeroEmpty(array)
}
func Logic04Soal04(n int) {
	array := array2.NewNumberArray(n, n) // untuk membuat array
	a := 1
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if i%4 == 0 { //untuk membuat deret angka dari kiri ke kanan
				array[i][n-1-j] = int32(a)
				a += 3
			} else if i%4 == 1 && j == 0 { // untuk membuat kotak kosong pada baris 1 dan berisi pada kolol 8
				array[i][j] = int32(a)
				a += 3
			} else if i%4 == 2 { // untuk membuat baris ke 3 dan 6 bergerak secara mundur
				array[i][j] = int32(a)
				a += 3
			} else if i%4 == 3 && j == n-1 { // untuk membuat isi pada kolom 0
				array[i][j] = int32(a)
				a += 3
			}
		}
	}
	array2.PrintNumberArrayZeroEmpty(array) // untuk print array2dimensi
}

func Logic04Soal05(n int) {
	array := array2.NewNumberArray(n, n)
	a := 1

	for i := 1; i <= n; i++ {
		for j := 1; j < n; j++ {
			if i == 1 {
				array[n-1-j][0] = int32(a)
			} else if i == 2 {
				array[0][j] = int32(a)
			} else if i == 3 {
				array[j][n-1] = int32(a)
			} else if i == 4 {
				array[n-1][n-1-j] = int32(a)
			}
			a++
		}
	}
	array2.PrintNumberArrayZeroEmpty(array)
}

func Logic04Soal06(n int) {
	array := array2.NewNumberArray(n, n)
	a := 1

	for i := 1; i <= n; i++ {
		for j := 1; j < n; j++ {
			if i == 1 {
				array[j][0] = int32(a)
			} else if i == 2 {
				array[n-1][j] = int32(a)
			} else if i == 3 {
				array[n-1-j][n-1] = int32(a)
			} else if i == 4 {
				array[0][n-1-j] = int32(a)
			}
			a++
		}
	}
	array2.PrintNumberArrayZeroEmpty(array)
}
