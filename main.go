package main

import "fmt"

// Basic deklarasi variabel

// func main() {
// 	var nama string = "Dimas"
// 	umur := 30
// 	tinggiBadan := 170.5
// 	sudahMenikah := true

// 	fmt.Println("Nama:", nama)
// 	fmt.Println("Umur:", umur)
// 	fmt.Println("Tinggi:", tinggiBadan, "cm")
// 	fmt.Println("Sudah Menikah:", sudahMenikah)
// }

// Kondisi

// func main() {
// 	nilai := 85

// 	if nilai >= 90 {
// 		fmt.Println("Grade A")
// 	} else if nilai >= 80 {
// 		fmt.Println("Grade B")
// 	} else if nilai >= 70 {
// 		fmt.Println("Grade C")
// 	} else {
// 		fmt.Println("Grade D")
// 	}
// }

// Perulangan Loop

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("Perulangan ke-", i)
// 	}
// }

// func main() {
// 	hitungMundur := 5

// 	for hitungMundur > 0 {
// 		fmt.Println(hitungMundur)
// 		hitungMundur--
// 	}
// 	fmt.Println("Selesai!")
// }

//  Slice Array

// func main() {
// 	buahBuahan := []string{"Apel", "Jeruk", "Mangga"}

// 	fmt.Println("Daftar buah:", buahBuahan)
// 	fmt.Println("Buah Pertama", buahBuahan[0])
// 	fmt.Println("Jumlah Buah", len(buahBuahan))

// 	buahBuahan = append(buahBuahan, "Pisang")
// 	fmt.Println("Setelah nambah:", buahBuahan)
// }

// func main() {
// 	buahBuahan := []string{"Apel", "Jeruk", "Mangga", "Pisang"}

// 	for i, buah := range buahBuahan {
// 		fmt.Println("Index", i, ":", buah)
// 	}
// }

//  Map golang

// func main() {
// 	biodata := map[string]string{
// 		"nama": "Dimas",
// 		"kota": "Balikpapan",
// 		"hobi": "ngoding",
// 	}

// 	fmt.Println("Nama", biodata["nama"])
// 	fmt.Println("Kota", biodata["kota"])

// 	biodata["pekerjaan"] = "Programmer"
// 	fmt.Println("Semua data:", biodata)
// }

// func main() {
// 	biodata := map[string]string{
// 		"nama": "Dimas",
// 		"kota": "Balikpapan",
// 		"hobi": "Ngoding",
// 	}

// 	for key, value := range biodata {
// 		fmt.Println(key, ":", value)
// 	}
// }

// GolangFunction

// func sapa(nama string) {
// 	fmt.Println("Halo,", nama, "!")
// }

// func main() {
// 	sapa("Dimas")
// 	sapa("Budi")
// }

// func tambah(a int, b int) int {
// 	hasil := a + b
// 	return hasil
// }

// func main() {
// 	total := tambah(5, 3)
// 	fmt.Println("Hasil:", total)
// }

// func bagiDua(angka int) (int, int) {
// 	hasil := angka / 2
// 	sisa := angka % 2
// 	return hasil, sisa
// }

// func main() {
// 	hasil, sisa := bagiDua(7)
// 	fmt.Println("Hasil bagi:", hasil)
// 	fmt.Println("Sisa:", sisa)
// }

// Error handling

// func bagi(a int, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("tidak bisa membagi dengan nol")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	hasil, err := bagi(2, 0)
// 	if err != nil {
// 		fmt.Println("Terjadi error:", err)
// 	} else {
// 		fmt.Println("Hasil:", hasil)
// 	}
// }

// Basic Struct

// type Orang struct {
// 	Nama string
// 	Umur int
// 	Kota string
// }

// func main() {
// 	orang1 := Orang {
// 		Nama: "Dimas",
// 		Umur: 30,
// 		Kota: "Balikpapan",
// 	}

// 	fmt.Println("Nama:", orang1.Nama)
// 	fmt.Println("Umur:", orang1.Umur)
// 	fmt.Println("Kota:", orang1.Kota)
// 	fmt.Println("Semua data:", orang1)
// }

// type Orang struct {
// 	Nama string
// 	Umur int
// 	Kota string
// }

// func main() {
// 	orang1 := Orang{Nama: "Dimas", Umur: 30, Kota: "Balikpapan"}
// 	orang2 := Orang{Nama: "Riyani", Umur: 29, Kota: "Samarinda"}

// 	daftarOrang := []Orang{orang1, orang2}

// 	for _, orang := range daftarOrang {
// 		fmt.Println(orang.Nama, "-", orang.Umur, "tahun -", orang.Kota)
// 	}
// }

//  Struct & Function

// type Orang struct {
// 	Nama string
// 	Umur int
// 	Kota string
// }

// func (o Orang) Sapa() {
// 	fmt.Println("Halo, nama saya", o.Nama, "dari", o.Kota)
// }

// func main() {
// 	orang1 := Orang{Nama: "Dimas", Umur: 30, Kota: "Balikpapan"}
// 	orang1.Sapa()
// }

// type Orang struct {
// 	Nama string
// 	Umur int
// 	Kota string
// }

// func (o Orang) SudahDewasa() bool {
// 	return o.Umur >= 18
// }

// func main() {
// 	orang1 := Orang{Nama: "Dimas", Umur: 30, Kota: "Balikpapan"}
// 	orang2 := Orang{Nama: "Riyani", Umur: 15, Kota: "Samarinda"}

// 	fmt.Println(orang1.Nama, "sudah dewasa:", orang1.SudahDewasa())
// 	fmt.Println(orang2.Nama, "sudah dewasa:", orang2.SudahDewasa())
// }

// Pointer

// func main() {
// 	angka := 10
// 	pointerAngka := &angka

// 	fmt.Println("Nilai angka:", angka)
// 	fmt.Println("Alamat memori angka:", pointerAngka)
// 	fmt.Println("Nilai yang ditunjuk pointer:", *pointerAngka)
// }

// func tambahSepuluh(angka int) {
// 	angka = angka + 10
// }

// func tambahSepuluhPointer(angka *int) {
// 	*angka = *angka + 10
// }

// func main() {
// 	nilai1 := 20
// 	tambahSepuluh(nilai1)
// 	fmt.Println("Setelah tambahSepuluh:", nilai1)

// 	nilai2 := 5
// 	tambahSepuluhPointer(&nilai2)
// 	fmt.Println("Setelah tambahSepuluhPointer:", nilai2)
// }

//  struct + pointer

// type Orang struct {
// 	Nama string
// 	Umur int
// }

// func (o *Orang) TambahUmur() {
// 	o.Umur = o.Umur +1
// }

// func main() {
// 	orang1 := Orang{Nama: "Dimas", Umur: 30}
// 	orang1.TambahUmur()
// 	fmt.Println(orang1.Nama, "sekarang umur:", orang1.Umur)
// }