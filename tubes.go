package main

import "fmt"

const NMAX = 100

type Sampah struct {
	jenis  string
	jumlah int
	metode string
}

type tabSampah [NMAX]Sampah

func main() {
	var data tabSampah
	var n int
	var pilihan int

	for {
		fmt.Println("\nMenu Utama:")
		fmt.Println("1. Manajemen Data Sampah")
		fmt.Println("2. Pencarian Data Sampah")
		fmt.Println("3. Pengurutan Data Sampah")
		fmt.Println("4. Cetak Semua Data")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			menuManajemen(&data, &n)
		case 2:
			menuPencarian(data, n)
		case 3:
			menuPengurutan(&data, n)
		case 4:
			cetakData(data, n)
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuManajemen(T *tabSampah, n *int) {
	var pilih int
	for {
		fmt.Println("\nMenu Manajemen Sampah:")
		fmt.Println("1. Tambah Data Sampah")
		fmt.Println("2. Ubah Data Sampah")
		fmt.Println("3. Hapus Data Sampah")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahData(T, n)
			fmt.Println("Data Sampah berhasil ditambahkan!")
		case 2:
			ubahData(T, *n)
		case 3:
			hapusData(T, n)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
func menuPencarian(T tabSampah, n int) {
	var pilih int
	for {
		fmt.Println("\nMenu Pencarian:")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search (harus diurutkan berdasar jenis dahulu)")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			seqSearch(T, n)
		case 2:
			binSearch(T, n)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
func menuPengurutan(T *tabSampah, n int) {
	var pilih int
	for {
		fmt.Println("\nMenu Pengurutan:")
		fmt.Println("1. Urutkan Berdasarkan Jumlah (Selection Sort)")
		fmt.Println("2. Urutkan Berdasarkan Jenis (Insertion Sort)")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			selectionSortJumlah(T, n)
			fmt.Println("Data berhasil diurutkan berdasarkan jumlah.")
		case 2:
			insertionSortJenis(T, n)
			fmt.Println("Data berhasil diurutkan berdasarkan jenis.")
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
func tambahData(T *tabSampah, n *int) {
	if *n < NMAX {
		fmt.Print("Jenis sampah: ")
		fmt.Scan(&T[*n].jenis)
		fmt.Print("Jumlah sampah (kg): ")
		fmt.Scan(&T[*n].jumlah)
		fmt.Print("Metode daur ulang: ")
		fmt.Scan(&T[*n].metode)
		*n++
	} else {
		fmt.Println("Kapasitas penuh!")
	}
}

func ubahData(T *tabSampah, n int) {
	var jenis, metode string
	var jumlah int
	var i int

	fmt.Print("Masukkan jenis sampah yang ingin diubah: ")
	fmt.Scan(&jenis)
	fmt.Print("Masukkan jumlah sampah yang ingin diubah: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan metode daur ulang yang ingin diubah: ")
	fmt.Scan(&metode)

	for i = 0; i < n; i++ {
		if T[i].jenis == jenis && T[i].jumlah == jumlah && T[i].metode == metode {
			fmt.Print("Jenis baru: ")
			fmt.Scan(&T[i].jenis)
			fmt.Print("Jumlah baru: ")
			fmt.Scan(&T[i].jumlah)
			fmt.Print("Metode baru: ")
			fmt.Scan(&T[i].metode)
			fmt.Println("Data berhasil diubah!")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}
func hapusData(T *tabSampah, n *int) {
	var jenis, metode string
	var jumlah int
	var i, j int

	fmt.Print("Masukkan jenis sampah yang ingin dihapus: ")
	fmt.Scan(&jenis)
	fmt.Print("Masukkan jumlah sampah yang ingin dihapus: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan metode daur ulang yang ingin dihapus: ")
	fmt.Scan(&metode)

	for i = 0; i < *n; i++ {
		if T[i].jenis == jenis && T[i].jumlah == jumlah && T[i].metode == metode {
			for j = i; j < *n-1; j++ {
				T[j] = T[j+1]
			}
			*n--
			fmt.Println("Data berhasil dihapus!")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}
func cetakData(T tabSampah, n int) {
	if n == 0 {
		fmt.Println("Belum ada data.")
	} else {
		fmt.Println("\nData Sampah:")
		for i := 0; i < n; i++ {
			fmt.Println(i+1, "-", T[i].jenis, ",", T[i].jumlah, "kg,", T[i].metode)
		}
	}
}
func seqSearch(T tabSampah, n int) {
	var jenis string
	var ditemukan bool = false

	fmt.Print("Masukkan jenis sampah yang dicari: ")
	fmt.Scan(&jenis)

	for i := 0; i < n; i++ {
		if T[i].jenis == jenis {
			fmt.Println("Ditemukan:", T[i].jenis, ",", T[i].jumlah, "kg,", T[i].metode)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Data tidak ditemukan!")
	}
}


func binSearch(T tabSampah, n int) {
	var jenis string
	fmt.Print("Masukkan jenis sampah yang dicari: ")
	fmt.Scan(&jenis)
	kiri := 0
	kanan := n - 1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if T[tengah].jenis == jenis {
			fmt.Println("Ditemukan:", T[tengah].jenis, ",", T[tengah].jumlah, "kg,", T[tengah].metode)
			return
		} else if T[tengah].jenis < jenis {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func selectionSortJumlah(T *tabSampah, n int) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if T[j].jumlah < T[idx].jumlah {
				idx = j
			}
		}
		T[i], T[idx] = T[idx], T[i]
	}
}

func insertionSortJenis(T *tabSampah, n int) {
	for i := 1; i < n; i++ {
		temp := T[i]
		j := i - 1
		for j >= 0 && T[j].jenis > temp.jenis {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = temp
	}
}