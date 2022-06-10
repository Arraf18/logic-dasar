package logic_01

import (
	"fmt"
	"math"
	"testing"
)

//soal 1
func TestSoal01(t *testing.T) {
	n := 10
	angka := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
		}
	}
}

//soal 2
func TestSoa0l2(t *testing.T) {
	n := 10
	angka := 1

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
		}
	}
}

//soal 3

func TestSoa0l3(t *testing.T) {
	n := 10
	x := 99
	angka := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(x, "\t")
		}
	}
}

//soal 4
func TestSoa0l4(t *testing.T) {
	n := 10
	x := 777
	angka := 1

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(x, "\t")
		}
	}
}

//soal 5
func TestSoal05(t *testing.T) {
	n := 15
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("fizz buzz", "\t")
		} else if i%3 == 0 {
			fmt.Print("fizz", "\t")
		} else if i%5 == 0 {
			fmt.Print("buzz", "\t")
		} else {
			fmt.Print(i, "\t")
		}
	}
}

func TestSoal06(t *testing.T) {
	n := 15
	x := 1
	y := 24
	z := 12
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			x = i * i
			fmt.Print(x, "\t")
		} else if i <= 7 {
			x = i * 3
			fmt.Print(x, "\t")
		} else if i <= 11 {
			y -= 3
			fmt.Print(y, "\t")
		} else {
			z -= 3
			fmt.Print(z, "\t")
		}
	}
}

//Cara Pak Roni

func TestPakRoni06(t *testing.T) {
	n := 15
	nt := n / 2
	angka := 3

	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			x := math.Pow(float64(i), 2)
			fmt.Print(x, "\t")
		} else {
			fmt.Print(angka, "\t")
		}

		if i <= nt {
			angka = angka + 3
		} else {
			angka = angka - 3

		}
	}
}

func TestSoal07(t *testing.T) {
	n := 10
	angka := 4
	x := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
		} else {
			x += 3
			fmt.Print(x, "\t")
		}
	}
}

func TestSoal08(t *testing.T) {
	n := 10
	a := 2
	b := 5
	for i := 1; i <= n; i += 1 {
		if i%2 == 0 {
			fmt.Print(b, "\t")
			b *= 5
		} else {
			fmt.Print(a, "\t")
			a *= 2
		}
	}
}

func TestSoal09(t *testing.T) {
	n := 12
	a := 1 //a bilangan yg akan dipangkatkan
	b := 1 //pangkat bilangan
	for i := 1; i <= n; i++ {
		fmt.Print(math.Pow(float64(a), float64(b)), "\t") //print angka pertama
		if i%3 == 0 {
			a = 1  //a reset kembali ke 1 di tiap index 3
			b += 1 // pangkat naik +1 tiap index 3 : 1,1,1 2,2,2 3,3,3
		} else {
			a += 1 // tambah angka +1 a : 2,3,, 1,2,3,, 1,2,3
		}
	}
}

func TestSoal10(t *testing.T) {
	n := 12
	a := 1 // bilangan yg akan dikalikan
	b := 4 // pengali bilangan
	for i := 1; i <= n; i++ {
		if i%4 == 0 { // kondisi tiap index ke 4 untuk merubah kondisi a ke tahap selanjutnya, dan b
			fmt.Print(999, "\t")
			a += 1 // a dinaikan +1 agar bilangan jadi 2
			b = 4  // b di reset ke 4 untuk dikurangi lagi agar penggali jadi 3,2,1
		} else {
			b -= 1
			fmt.Print(a*b, "\t")
		}
	}
}
