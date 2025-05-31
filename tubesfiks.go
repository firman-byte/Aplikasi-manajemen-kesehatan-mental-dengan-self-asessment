package main

import "fmt"

const NMAX int = 100

type Pengguna struct {
	nama        string
	umur        int
	totalSkor   int
	jumlahData  int
	riwayatSkor SkorArray
}

type SkorArray [NMAX]int

func main() {
	var user Pengguna
	var pilihan int

	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&user.nama)
	fmt.Print("Masukkan Umur : ")
	fmt.Scan(&user.umur)

	fmt.Println("\nSelamat datang,", user.nama)

	for {
		fmt.Println("\n=== Menu Self Assessment ===")
		fmt.Println("1. Lakukan Self Assessment")
		fmt.Println("2. Hasil Skor Self Assessment")
		fmt.Println("3. Lihat Riwayat")
		fmt.Println("4. Saran Penenangan")
		fmt.Println("5. Urutkan Data Hasil Skor (Selection Sort)")
		fmt.Println("6. Urutkan Data Hasil Skor (Insertion Sort)")
		fmt.Println("7. Cari Skor (Sequential Search)")
		fmt.Println("8. Cari Skor (Binary Search)")
		fmt.Println("9. Pencarian Berdasarkan Kategori Tingkat Stres")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu (0-9): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 0:
			fmt.Println("Terima kasih telah menggunakan self-assessment ini.")
			return
		case 1:
			selfAssessment(&user)
			totalSkorResult(user.totalSkor, user.umur)
		case 2:
			totalSkorResult(user.totalSkor, user.umur)
		case 3:
			lihatRiwayat(user.riwayatSkor, user.jumlahData)
		case 4:
			menuSaranPenenangan()
		case 5:
			selectionSort(&user.riwayatSkor, user.jumlahData)
			lihatRiwayat(user.riwayatSkor, user.jumlahData)
		case 6:
			insertionSort(&user.riwayatSkor, user.jumlahData)
			lihatRiwayat(user.riwayatSkor, user.jumlahData)
		case 7:
			sequentialSearch(user.riwayatSkor, user.jumlahData)
		case 8:
			binarySearch(user.riwayatSkor, user.jumlahData)
		case 9:
			cariKategori(user.riwayatSkor, user.jumlahData, user.umur)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func selfAssessment(user *Pengguna) {
	var jawaban int
	var i int
	var pertanyaan [10]string

	pertanyaan[0] = "1. Saya sering merasa tegang atau gelisah: "
	pertanyaan[1] = "2. Saya sulit untuk rileks: "
	pertanyaan[2] = "3. Saya merasa panik tanpa alasan yang jelas: "
	pertanyaan[3] = "4. Saya terlalu khawatir berlebihan terhadap hal hal kecil: "
	pertanyaan[4] = "5. Saya merasa lelah secara mental: "
	pertanyaan[5] = "6. Saya merasa kehilangan motivasi hidup : "
	pertanyaan[6] = "7. Saya sulit tidur atau tidur saya tidak nyenyak: "
	pertanyaan[7] = "8. Saya merasa jantung berdebar tanpa sebab yang jelas: "
	pertanyaan[8] = "9. Saya mudah tersinggung atau marah: "
	pertanyaan[9] = "10. Saya merasa tidak bisa mengontrol kekhawatiran saya: "

	user.totalSkor = 0
	fmt.Println("Jawab dengan skala (0 = Tidak Pernah, 1 = Kadang-kadang, 2 = Sering, 3 = Sangat Sering)")
	for i = 0; i < 10; i++ {
		fmt.Print(pertanyaan[i])
		fmt.Scan(&jawaban)
		for jawaban < 0 || jawaban > 3 {
			fmt.Print("Masukkan hanya 0-3. Coba lagi: ")
			fmt.Scan(&jawaban)
		}
		user.totalSkor += jawaban
	}

	if user.jumlahData < NMAX {
		user.riwayatSkor[user.jumlahData] = user.totalSkor
		user.jumlahData++
	} else {
		fmt.Println("Riwayat skor sudah penuh.")
	}

	fmt.Println("\nJawaban anda telah tersimpan. Terima kasih.")
}

func totalSkorResult(n int, umur int) {
	fmt.Println("Total Skor Anda:", n)
	if umur >= 13 && umur <= 17 {
		if n <= 7 {
			fmt.Println("Tingkat stres : Rendah")
		} else if n <= 14 {
			fmt.Println("Tingkat stres : Sedang")
		} else if n <= 20 {
			fmt.Println("Tingkat stres : Tinggi")
		} else {
			fmt.Println("Tingkat stres : Sangat Tinggi")
		}
	} else if umur >= 18 && umur <= 25 {
		if n <= 10 {
			fmt.Println("Tingkat stres : Rendah")
		} else if n <= 17 {
			fmt.Println("Tingkat stres : Sedang")
		} else if n <= 22 {
			fmt.Println("Tingkat stres : Tinggi")
		} else {
			fmt.Println("Tingkat stres : Sangat Tinggi")
		}
	} else {
		if n <= 12 {
			fmt.Println("Tingkat stres : Rendah")
		} else if n <= 19 {
			fmt.Println("Tingkat stres : Sedang")
		} else if n <= 24 {
			fmt.Println("Tingkat stres : Tinggi")
		} else {
			fmt.Println("Tingkat stres : Sangat Tinggi")
		}
	}
}

func lihatRiwayat(data SkorArray, n int) {
	if n == 0 {
		fmt.Println("Belum ada riwayat skor.")
	} else {
		fmt.Println("\n--- Riwayat Skor Self Assessment ---")
		for i := 0; i < n; i++ {
			fmt.Printf("%d. Skor: %d\n", i+1, data[i])
		}
	}
}

func selectionSort(data *SkorArray, n int) {
	var i, j, idx_max int
	for i = 0; i < n-1; i++ {
		idx_max = i
		for j = i + 1; j < n; j++ {
			if data[j] > data[idx_max] {
				idx_max = j
			}
		}
		data[i], data[idx_max] = data[idx_max], data[i]
	}
}

func insertionSort(data *SkorArray, n int) {
	var i, j, temp int
	for i = 1; i < n; i++ {
		temp = data[i]
		j = i
		for j > 0 && temp < data[j-1] {
			data[j] = data[j-1]
			j--
		}
		data[j] = temp
	}
}

func sequentialSearch(data SkorArray, n int) {
	var cari, i int
	fmt.Print("Masukkan skor yang ingin dicari: ")
	fmt.Scan(&cari)
	for i = 0; i < n; i++ {
		if data[i] == cari {
			fmt.Printf("Skor %d ditemukan pada riwayat ke-%d\n", cari, i+1)
			return
		}
	}
	fmt.Println("Skor tidak ditemukan dalam riwayat.")
}

func binarySearch(data SkorArray, n int) {
	var kiri, kanan, tengah, x int
	fmt.Print("Masukkan skor yang ingin dicari: ")
	fmt.Scan(&x)
	kiri = 0
	kanan = n - 1
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if data[tengah] == x {
			fmt.Printf("Skor %d ditemukan pada indeks ke-%d\n", x, tengah+1)
			return
		} else if x < data[tengah] {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	fmt.Println("Skor tidak ditemukan.")
}

func cariKategori(data SkorArray, n int, umur int) {
	var kategori string
	var i int
	fmt.Print("Masukkan kategori (rendah/sedang/tinggi/sangat tinggi): ")
	fmt.Scan(&kategori)
	for i = 0; i < n; i++ {
		if cocokKategori(data[i], umur, kategori) {
			fmt.Printf("%d. Skor: %d\n", i+1, data[i])
		}
	}
}

func cocokKategori(skor int, umur int, kategori string) bool {
	if umur >= 13 && umur <= 17 {
		switch kategori {
		case "rendah":
			return skor <= 7
		case "sedang":
			return skor > 7 && skor <= 14
		case "tinggi":
			return skor > 14 && skor <= 20
		case "sangat tinggi":
			return skor > 20
		}
	} else if umur >= 18 && umur <= 25 {
		switch kategori {
		case "rendah":
			return skor <= 10
		case "sedang":
			return skor > 10 && skor <= 17
		case "tinggi":
			return skor > 17 && skor <= 22
		case "sangat tinggi":
			return skor > 22
		}
	} else {
		switch kategori {
		case "rendah":
			return skor <= 12
		case "sedang":
			return skor > 12 && skor <= 19
		case "tinggi":
			return skor > 19 && skor <= 24
		case "sangat tinggi":
			return skor > 24
		}
	}
	return false
}

func menuSaranPenenangan() {
	var pilih int
	for {
		fmt.Println("\n=== Menu Saran Penenangan ===")
		fmt.Println("1. Rekomendasi Musik")
		fmt.Println("2. Cerita Tidur")
		fmt.Println("3. Puisi Tidur")
		fmt.Println("0. Kembali ke menu utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 0:
			return
		case 1:
			rekomendasiMusik()
		case 2:
			ceritaTidur()
		case 3:
			puisiTidur()
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func rekomendasiMusik() {
	var kembali int
	for {
		fmt.Println("\n--- Rekomendasi Musik ---")
		fmt.Println("1. Suara air hujan")
		fmt.Println("2. Suara sungai")
		fmt.Println("3. Suara hutan di malam hari")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&kembali)
		switch kembali {
		case 0:
			return
		case 1:
			fmt.Println("Memulai suara air hujan...")
		case 2:
			fmt.Println("Memulai suara sungai...")
		case 3:
			fmt.Println("Memulai suara hutan di malam hari...")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func ceritaTidur() {
	var pilih int
	for {
		fmt.Println("\n--- Cerita Tidur ---")
		fmt.Println("1. Dongeng")
		fmt.Println("2. Kisah-kisah nabi")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)
		switch pilih {
		case 0:
			return
		case 1:
			dongengTidur()
		case 2:
			kisahNabi()
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func dongengTidur() {
	fmt.Println("\n--- Pilihan Dongeng ---")
	fmt.Println("1. Si Kancil dan Buaya")
	fmt.Println("2. Timun Mas")
	fmt.Println("3. Bawang Merah dan Bawang Putih")
	fmt.Print("Pilih dongeng (1-3): ")
	var pilih int
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		fmt.Println("Cerita: Si Kancil menipu buaya agar bisa menyeberangi sungai")
	case 2:
		fmt.Println("Cerita: Timun Mas lari dari kejaran raksasa dengan bantuan benda ajaib")
	case 3:
		fmt.Println("Cerita: Bawang Putih yang sabar akhirnya mendapatkan kebahagiaan")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func kisahNabi() {
	fmt.Println("\n--- Kisah Nabi ---")
	fmt.Println("1. Nabi Adam")
	fmt.Println("2. Nabi Nuh")
	fmt.Println("3. Nabi Ibrahim")
	fmt.Println("4. Nabi Musa")
	fmt.Println("5. Nabi Isa")
	fmt.Println("6. Nabi Muhammad")
	fmt.Print("Pilih nabi (1-6): ")
	var pilih int
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		fmt.Println("Nabi Adam adalah manusia pertama yang diciptakan oleh Allah.")
	case 2:
		fmt.Println("Nabi Nuh membuat kapal besar untuk menyelamatkan pengikutnya dari banjir besar.")
	case 3:
		fmt.Println("Nabi Ibrahim adalah bapak para nabi dan terkenal dengan kisah pengorbanannya.")
	case 4:
		fmt.Println("Nabi Musa membelah Laut Merah untuk menyelamatkan Bani Israil.")
	case 5:
		fmt.Println("Nabi Isa lahir dari Maryam dan memiliki banyak mukjizat.")
	case 6:
		fmt.Println("Nabi Muhammad adalah nabi terakhir yang membawa ajaran Islam.")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func puisiTidur() {
	fmt.Println("\n--- Puisi Tidur ---")
	fmt.Println("1. Senja yang Tenang")
	fmt.Println("2. Malam yang Damai")
	fmt.Println("3. Hening Sang Fajar")
	fmt.Print("Pilih puisi (1-3): ")
	var pilih int
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		fmt.Println("Puisi: Senja hadir membawa damai, perlahan gelap datang merayap, pejamkan matamu, dan beristirahatlah...")
	case 2:
		fmt.Println("Puisi: Di malam yang damai, bintang bercerita sunyi, selimuti dirimu dalam tenang malam...")
	case 3:
		fmt.Println("Puisi: Saat fajar menyingsing, sepi membungkus dunia, terlelaplah dalam damainya pagi...")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}