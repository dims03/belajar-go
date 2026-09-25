package main

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
