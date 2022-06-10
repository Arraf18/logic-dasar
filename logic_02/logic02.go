package logic_02

import "fmt"

//type data yg dipesan oleh function
func Logic02Soal01(n int) {
	a := 3

	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		}
		//kolom selesai
		//kemudian ke baris sebelumnya
		fmt.Println("\n")
		a += 3
	}
}

func Logic02Soal02(n int) {
	//looping baris
	for i := 0; i < n; i++ {
		a := 3

		//looping kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
			a += 3
		}
		//pindah baris
		fmt.Println("\n")
	}
}

func Logic02Soal03(n int) {
	//looping baris
	for i := 0; i < n; i++ {
		a := 3 * n
		//looping kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
			a -= 3
		}
		//pindah baris
		fmt.Println("\n")
	}
}

func Logic02Soal04(n int) {
	a := 27
	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		}
		//kolom selesai
		//pindah ke baris selanjutnya
		fmt.Println("\n")
	}
}

func Logic02Soal05(n int) {
	a := 3
	//ini nilai tengah
	nt := n / 2

	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		}
		//kolom selesai
		//kemudian pindah baris selanjutnya
		fmt.Println("\n")
		if i < nt {
			a += 3
			fmt.Print(a, "\t")
		} else {
			a -= 3
			fmt.Print(a, "\t")
		}
	}
}

func Logic02Soal06(n int) {
	nt := n / 2
	//looping baris
	for i := 0; i < n; i++ {
		a := 3

		//looping kolom
		for j := 0; j < n; j++ {
			if j < nt {
				fmt.Print(a, "\t")
				a += 3
			} else {
				fmt.Print(a, "\t")
				a -= 3
			}
		}
		//pindah baris
		fmt.Print("\n")
	}
}

func Logic02Soal07(n int) {
	//ini variabel a, nilainya 3
	a := 3
	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			if i < j {
				fmt.Print(" ", "\t")
			} else {
				//ini menampilkan kolom
				fmt.Print(a, "\t")
			}
		}
		//kolom selesai
		//kemudain pindah ke baris selanjutnya
		fmt.Print("\n")
		//saat pindah baris variabel a ditambah 3
		a += 3
	}
}

func Logic02Soal08(n int) {
	a := 3
	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			//membuat kondisi untuk kotak yg kosong
			if i > j {
				fmt.Print(" ", "\t")
			} else {
				fmt.Print(a, "\t")
			}
		}
		//kolom selesai
		//pindah kenbaris selanjutnya
		fmt.Print("\n")
		//nilai variabel a ditambah 3
		a += 3
	}
}

func Logic02Soal09(n int) {
	for i := 0; i < n; i++ {
		a := 0
		for j := 0; j < n; j++ {
			a += 3
			if i+j == n-1 {
				fmt.Print(a, "\t")
			} else if i+j <= n-1 {
				fmt.Print(" ", "\t")
			} else {
				fmt.Print(a, "\t")
			}
		}
		fmt.Print("\n")
	}
}

func Logic02Soal10(n int) {
	for i := 0; i < n; i++ {
		a := 0
		for j := 0; j < n; j++ {
			a += 3
			if i+j == n-1 {
				fmt.Print(a, "\t")
			} else if i+j <= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Print("\n")
	}
}
