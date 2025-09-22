package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var choice int
		fmt.Println("=================================================")
		fmt.Println("=== Selamat Datang di Kalkulator Bangun Datar ===")
		fmt.Println("===            1. Lingkaran                   ===")
		fmt.Println("===            2. Persegi                     ===")
		fmt.Println("===            3. Balok                       ===")
		fmt.Println("===            4. Layang-Layang               ===")
		fmt.Println("===            0. Exit                        ===")
		fmt.Println("=================================================")
		fmt.Print("Silahkan memilih angka tersebut : ")
		fmt.Scan(&choice)
		fmt.Println("\n\n=================================================")

		switch choice {
		case 1:
			lingkaran()
		case 2:
			persegi()
		case 3:
			balok()
		case 4:
			layangLayang()
		case 0:
			fmt.Println("Program berakhir sampai bertemu lagi, karena kamu akan selalu ada di hatiku ❤️")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
		fmt.Println("")
	}
}

func lingkaran() {
	var choice int
	fmt.Println("=================================================")
	fmt.Println("===       1. Hitung keliling lingkaran        ===")
	fmt.Println("===       2. Hitung Luas lingkaran            ===")
	fmt.Println("===       0. Exit                             ===")
	fmt.Print("Silahkan pilih : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		hitungKelilingLingkaran()
	case 2:
		hitungLingkaran()
	case 0:
		return
	default:
		fmt.Println("Silahkan memilih dengan benar")
	}
}

func persegi() {
	var choice int
	fmt.Println("=================================================")
	fmt.Println("===       1. Hitung keliling Persegi          ===")
	fmt.Println("===       2. Hitung Luas Persegi              ===")
	fmt.Println("===       0. Exit                             ===")
	fmt.Print("Silahkan pilih : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		hitungKelilingPersegi()
	case 2:
		hitungLuasPersegi()
	case 0:
		return
	default:
		fmt.Println("Silahkan memilih dengan benar")
	}
}

func balok() {
	var choice int
	fmt.Println("=================================================")
	fmt.Println("===       1. Hitung Keliling Balok            ===")
	fmt.Println("===       2. Hitung Luas Permukaan Balok      ===")
	fmt.Println("===       0. Exit                             ===")
	fmt.Print("Silahkan pilih : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		hitungKelilingBalok()
	case 2:
		hitungLuasBalok()
	case 0:
		return
	default:
		fmt.Println("Silahkan memilih dengan benar")
	}
}

func layangLayang() {
	var choice int
	fmt.Println("=================================================")
	fmt.Println("===       1. Hitung Keliling Layang-Layang    ===")
	fmt.Println("===       2. Hitung Luas Layang-Layang        ===")
	fmt.Println("===       0. Exit                             ===")
	fmt.Print("Silahkan pilih : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		hitungKelilingLayang()
	case 2:
		hitungLuasLayang()
	case 0:
		return
	default:
		fmt.Println("Silahkan memilih dengan benar")
	}
}

// ================= PERSEGI =================
func hitungKelilingPersegi() {
	var s float64
	fmt.Print("Masukan sisi : ")
	fmt.Scan(&s)
	result := kelilingPersegi(s)
	fmt.Println("Jadi keliling persegi dengan sisi ", s, " adalah : ", result)
}

func kelilingPersegi(s float64) float64 {
	return 4 * s
}

func hitungLuasPersegi() {
	var s float64
	fmt.Print("Masukan sisi : ")
	fmt.Scan(&s)
	result := luasPersegi(s)
	fmt.Println("Jadi luas persegi dengan sisi ", s, " adalah : ", result)
}

func luasPersegi(s float64) float64 {
	return s * s
}

// ================= LINGKARAN =================
func hitungLingkaran() {
	var r float64
	fmt.Print("Masukan jari-jari dalam cm : ")
	fmt.Scan(&r)

	if r <= 0 {
		fmt.Println("Tolong masukan jari-jari dengan benar dan harus lebih dari 0 :) ")
		return
	}
	result := luasLingkaran(r)
	fmt.Println("Luas lingkaran dari jari-jari ", r, " adalah : ", result)
}

func hitungKelilingLingkaran() {
	var d float64
	fmt.Print("Masukan diameter lingkaran : ")
	fmt.Scan(&d)

	if d <= 0 {
		fmt.Println("Tolong masukan diameter dengan benar dan harus lebih dari 0")
		return
	}

	result := kelilingLingkaran(d)
	fmt.Println("Keliling lingkaran dari diameter ", d, " adalah : ", result)
}

func luasLingkaran(r float64) float64 {
	piValue := math.Pi
	return piValue * r * r
}

func kelilingLingkaran(d float64) float64 {
	piValue := math.Pi
	return piValue * d
}

// ================= BALOK =================
func hitungLuasBalok() {
	var p, l, t float64
	fmt.Print("Masukan panjang : ")
	fmt.Scan(&p)
	fmt.Print("Masukan lebar : ")
	fmt.Scan(&l)
	fmt.Print("Masukan tinggi : ")
	fmt.Scan(&t)

	result := luasBalok(p, l, t)
	fmt.Println("Luas permukaan balok adalah : ", result)
}

func hitungKelilingBalok() {
	var p, l, t float64
	fmt.Print("Masukan panjang : ")
	fmt.Scan(&p)
	fmt.Print("Masukan lebar : ")
	fmt.Scan(&l)
	fmt.Print("Masukan tinggi : ")
	fmt.Scan(&t)

	result := kelilingBalok(p, l, t)
	fmt.Println("Keliling (jumlah panjang semua rusuk) balok adalah : ", result)
}

func luasBalok(p, l, t float64) float64 {
	return 2 * (p*l + p*t + l*t)
}

func kelilingBalok(p, l, t float64) float64 {
	return 4 * (p + l + t)
}

// ================= LAYANG-LAYANG =================
func hitungLuasLayang() {
	var d1, d2 float64
	fmt.Print("Masukan diagonal 1 : ")
	fmt.Scan(&d1)
	fmt.Print("Masukan diagonal 2 : ")
	fmt.Scan(&d2)

	result := luasLayang(d1, d2)
	fmt.Println("Luas layang-layang adalah : ", result)
}

func hitungKelilingLayang() {
	var a, b float64
	fmt.Print("Masukan sisi panjang : ")
	fmt.Scan(&a)
	fmt.Print("Masukan sisi pendek : ")
	fmt.Scan(&b)

	result := kelilingLayang(a, b)
	fmt.Println("Keliling layang-layang adalah : ", result)
}

func luasLayang(d1, d2 float64) float64 {
	return (d1 * d2) / 2
}

func kelilingLayang(a, b float64) float64 {
	return 2 * (a + b)
}
