package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const MAX_SUBSCRIPTIONS = 100

type Tanggal struct {
	Hari  int
	Bulan int
	Tahun int
}

type Langganan struct {
	ID               int
	Nama             string
	Kategori         string
	BiayaBulanan     float64
	TanggalBayar     Tanggal
	MetodePembayaran string
	Status           string // "Aktif", "Tidak Aktif", "Akan Berakhir"
}

type DaftarLangganan struct {
	Data  [MAX_SUBSCRIPTIONS]Langganan
	Count int
}

var daftarLangganan DaftarLangganan

func tampilkanMenu() {
	fmt.Println("\n=== APLIKASI MANAJEMEN SUBSKRIPSI DIGITAL ===")
	fmt.Println("1. Tambah Langganan")
	fmt.Println("2. Lihat Semua Langganan")
	fmt.Println("3. Cari Langganan")
	fmt.Println("4. Edit Langganan")
	fmt.Println("5. Hapus Langganan")
	fmt.Println("6. Urutkan Langganan")
	fmt.Println("7. Laporan Keuangan")
	fmt.Println("8. Pengingat Pembayaran")
	fmt.Println("9. Rekomendasi Penghematan")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

func bacaString(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func bacaInt(prompt string) int {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	nilai, err := strconv.Atoi(input)
	if err != nil {
		return -1 // Mengembalikan -1 jik input tidak valid
	}
	return nilai
}

func bacaFloat(prompt string) float64 {
	input := bacaString(prompt)
	nilai, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0.0
	}
	return nilai
}

func bacaTanggal(prompt string) Tanggal {
	fmt.Print(prompt)
	var tanggal Tanggal
	fmt.Print("Hari (1-31): ")
	fmt.Scanf("%d", &tanggal.Hari)
	fmt.Print("Bulan (1-12): ")
	fmt.Scanf("%d", &tanggal.Bulan)
	fmt.Print("Tahun: ")
	fmt.Scanf("%d", &tanggal.Tahun)
	return tanggal
}

func tambahLangganan() {
	if daftarLangganan.Count >= MAX_SUBSCRIPTIONS {
		fmt.Println("Kapasitas penuh! Tidak dapat menambah langganan baru.")
		return
	}

	var langganan Langganan

	langganan.ID = daftarLangganan.Count + 1

	langganan.Nama = bacaString("Nama layanan: ")
	langganan.Kategori = bacaString("Kategori (Entertainment/Productivity/Education/Other): ")
	langganan.BiayaBulanan = bacaFloat("Biaya bulanan (Rp): ")
	langganan.TanggalBayar = bacaTanggal("Tanggal pembayaran:\n")
	langganan.MetodePembayaran = bacaString("Metode pembayaran: ")
	langganan.Status = "Aktif"

	daftarLangganan.Data[daftarLangganan.Count] = langganan
	daftarLangganan.Count++

	fmt.Println("Langganan berhasil ditambahkan!")
}

func tampilkanLangganan(langganan Langganan) {
	fmt.Printf("ID: %d\n", langganan.ID)
	fmt.Printf("Nama: %s\n", langganan.Nama)
	fmt.Printf("Kategori: %s\n", langganan.Kategori)
	fmt.Printf("Biaya Bulanan: Rp %.2f\n", langganan.BiayaBulanan)
	fmt.Printf("Tanggal Bayar: %d/%d/%d\n", langganan.TanggalBayar.Hari, langganan.TanggalBayar.Bulan, langganan.TanggalBayar.Tahun)
	fmt.Printf("Metode Pembayaran: %s\n", langganan.MetodePembayaran)
	fmt.Printf("Status: %s\n", langganan.Status)
	fmt.Println("---")
}

func lihatSemuaLangganan() {
	if daftarLangganan.Count == 0 {
		fmt.Println("Tidak ada langganan yang tersimpan.")
		return
	}

	fmt.Println("\n=== DAFTAR SEMUA LANGGANAN ===")
	for i := 0; i < daftarLangganan.Count; i++ {
		tampilkanLangganan(daftarLangganan.Data[i])
	}
}

func sequentialSearchNama(nama string) int {
	namaLower := strings.ToLower(nama)
	for i := 0; i < daftarLangganan.Count; i++ {
		if strings.Contains(strings.ToLower(daftarLangganan.Data[i].Nama), namaLower) {
			return i
		}
	}
	return -1
}

func sequentialSearchKategori(kategori string) []int {
	var hasil []int
	kategoriLower := strings.ToLower(kategori)

	for i := 0; i < daftarLangganan.Count; i++ {
		if strings.Contains(strings.ToLower(daftarLangganan.Data[i].Kategori), kategoriLower) {
			hasil = append(hasil, i)
		}
	}
	return hasil
}

func urutkanBerdasarkanNama(data []Langganan, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if strings.ToLower(data[j].Nama) < strings.ToLower(data[minIdx].Nama) {
				minIdx = j
			}
		}
		if minIdx != i {
			data[i], data[minIdx] = data[minIdx], data[i]
		}
	}
}

func binarySearchNama(nama string, data []Langganan, n int) int {
	namaLower := strings.ToLower(nama)
	left, right := 0, n-1

	for left <= right {
		mid := (left + right) / 2
		midNama := strings.ToLower(data[mid].Nama)

		if midNama == namaLower {
			return mid
		} else if midNama < namaLower {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func cariLangganan() {
	if daftarLangganan.Count == 0 {
		fmt.Println("Tidak ada langganan yang tersimpan.")
		return
	}

	fmt.Println("\n=== PENCARIAN LANGGANAN ===")
	fmt.Println("1. Cari berdasarkan nama (Sequential Search)")
	fmt.Println("2. Cari berdasarkan kategori (Sequential Search)")
	fmt.Println("3. Cari berdasarkan nama (Binary Search)")
	pilihan := bacaInt("Pilih metode pencarian: ")

	switch pilihan {
	case 1:
		nama := bacaString("Masukkan nama layanan: ")
		index := sequentialSearchNama(nama)
		if index != -1 {
			fmt.Println("\nLangganan ditemukan:")
			tampilkanLangganan(daftarLangganan.Data[index])
		} else {
			fmt.Println("Langganan tidak ditemukan.")
		}
	case 2:
		kategori := bacaString("Masukkan kategori: ")
		indices := sequentialSearchKategori(kategori)
		if len(indices) > 0 {
			fmt.Printf("\nDitemukan %d langganan dengan kategori '%s':\n", len(indices), kategori)
			for _, idx := range indices {
				tampilkanLangganan(daftarLangganan.Data[idx])
			}
		} else {
			fmt.Println("Tidak ada langganan dengan kategori tersebut.")
		}
	case 3:
		var dataTerurut [MAX_SUBSCRIPTIONS]Langganan
		for i := 0; i < daftarLangganan.Count; i++ {
			dataTerurut[i] = daftarLangganan.Data[i]
		}

		urutkanBerdasarkanNama(dataTerurut[:], daftarLangganan.Count)

		nama := bacaString("Masukkan nama layanan (harus tepat): ")
		index := binarySearchNama(nama, dataTerurut[:], daftarLangganan.Count)

		if index != -1 {
			fmt.Println("\nLangganan ditemukan:")
			tampilkanLangganan(dataTerurut[index])
		} else {
			fmt.Println("Langganan tidak ditemukan.")
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func cariIndexBerdasarkanID(id int) int {
	for i := 0; i < daftarLangganan.Count; i++ {
		if daftarLangganan.Data[i].ID == id {
			return i
		}
	}
	return -1
}

func editLangganan() {
	if daftarLangganan.Count == 0 {
		fmt.Println("Tidak ada langganan yang tersimpan.")
		return
	}

	lihatSemuaLangganan()
	id := bacaInt("Masukkan ID langganan yang akan diedit: ")

	if id <= 0 {
		fmt.Println("ID tidak valid. Silakan coba lagi.")
		return
	}

	index := cariIndexBerdasarkanID(id)
	if index == -1 {
		fmt.Println("Langganan dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println("\nData saat ini:")
	tampilkanLangganan(daftarLangganan.Data[index])

	fmt.Println("\nMasukkan data baru (tekan Enter untuk tidak mengubah):")

	nama := bacaString("Nama layanan: ")
	if nama != "" {
		daftarLangganan.Data[index].Nama = nama
	}

	kategori := bacaString("Kategori: ")
	if kategori != "" {
		daftarLangganan.Data[index].Kategori = kategori
	}

	biayaStr := bacaString("Biaya bulanan (Rp): ")
	if biayaStr != "" {
		biaya, err := strconv.ParseFloat(biayaStr, 64)
		if err == nil {
			daftarLangganan.Data[index].BiayaBulanan = biaya
		} else {
			fmt.Println("Biaya bulanan tidak valid. Data tidak diubah.")
		}
	}

	status := bacaString("Status (Aktif/Tidak Aktif/Akan Berakhir): ")
	if status != "" {
		daftarLangganan.Data[index].Status = status
	}

	fmt.Println("Langganan berhasil diperbarui!")
}

func hapusLangganan() {
	if daftarLangganan.Count == 0 {
		fmt.Println("Tidak ada langganan yang tersimpan.")
		return
	}

	lihatSemuaLangganan()
	id := bacaInt("Masukkan ID langganan yang akan dihapus: ")

	if id <= 0 {
		fmt.Println("ID tidak valid. Silakan coba lagi.")
		return
	}

	index := cariIndexBerdasarkanID(id)
	if index == -1 {
		fmt.Println("Langganan dengan ID tersebut tidak ditemukan.")
		return
	}

	// delete langganan
	for i := index; i < daftarLangganan.Count-1; i++ {
		daftarLangganan.Data[i] = daftarLangganan.Data[i+1]
	}
	daftarLangganan.Count--

	fmt.Println("Langganan berhasil dihapus!")
}

func selectionSortBiaya(ascending bool) {
	for i := 0; i < daftarLangganan.Count-1; i++ {
		extremeIdx := i
		for j := i + 1; j < daftarLangganan.Count; j++ {
			if ascending {
				if daftarLangganan.Data[j].BiayaBulanan < daftarLangganan.Data[extremeIdx].BiayaBulanan {
					extremeIdx = j
				}
			} else {
				if daftarLangganan.Data[j].BiayaBulanan > daftarLangganan.Data[extremeIdx].BiayaBulanan {
					extremeIdx = j
				}
			}
		}
		if extremeIdx != i {
			daftarLangganan.Data[i], daftarLangganan.Data[extremeIdx] = daftarLangganan.Data[extremeIdx], daftarLangganan.Data[i]
		}
	}
}

func bandingkanTanggal(t1, t2 Tanggal) int {
	if t1.Tahun != t2.Tahun {
		return t1.Tahun - t2.Tahun
	}
	if t1.Bulan != t2.Bulan {
		return t1.Bulan - t2.Bulan
	}
	return t1.Hari - t2.Hari
}

func insertionSortTanggal(ascending bool) {
	for i := 1; i < daftarLangganan.Count; i++ {
		key := daftarLangganan.Data[i]
		j := i - 1

		for j >= 0 {
			comparison := bandingkanTanggal(daftarLangganan.Data[j].TanggalBayar, key.TanggalBayar)
			shouldSwap := false

			if ascending && comparison > 0 {
				shouldSwap = true
			} else if !ascending && comparison < 0 {
				shouldSwap = true
			}

			if shouldSwap {
				daftarLangganan.Data[j+1] = daftarLangganan.Data[j]
				j--
			} else {
				break
			}
		}
		daftarLangganan.Data[j+1] = key
	}
}

func urutkanLangganan() {
	if daftarLangganan.Count == 0 {
		fmt.Println("Tidak ada langganan yang tersimpan.")
		return
	}

	fmt.Println("\n=== PENGURUTAN LANGGANAN ===")
	fmt.Println("1. Urutkan berdasarkan biaya (Selection Sort)")
	fmt.Println("2. Urutkan berdasarkan tanggal jatuh tempo (Insertion Sort)")
	pilihan := bacaInt("Pilih metode pengurutan: ")

	if pilihan == 1 || pilihan == 2 {
		fmt.Println("1. Urutan naik (ascending)")
		fmt.Println("2. Urutan turun (descending)")
		urutan := bacaInt("Pilih urutan: ")
		ascending := urutan == 1

		switch pilihan {
		case 1:
			selectionSortBiaya(ascending)
			fmt.Println("Langganan berhasil diurutkan berdasarkan biaya!")
		case 2:
			insertionSortTanggal(ascending)
			fmt.Println("Langganan berhasil diurutkan berdasarkan tanggal jatuh tempo!")
		}

		lihatSemuaLangganan()
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func hitungTotalPengeluaran() float64 {
	var total float64
	for i := 0; i < daftarLangganan.Count; i++ {
		if daftarLangganan.Data[i].Status == "Aktif" {
			total += daftarLangganan.Data[i].BiayaBulanan
		}
	}
	return total
}

func laporanKeuangan() {
	if daftarLangganan.Count == 0 {
		fmt.Println("Tidak ada langganan yang tersimpan.")
		return
	}

	fmt.Println("\n=== LAPORAN KEUANGAN ===")

	total := hitungTotalPengeluaran()
	fmt.Printf("Total pengeluaran bulanan: Rp %.2f\n", total)

	kategoriMap := make(map[string]float64)
	for i := 0; i < daftarLangganan.Count; i++ {
		if daftarLangganan.Data[i].Status == "Aktif" {
			kategoriMap[daftarLangganan.Data[i].Kategori] += daftarLangganan.Data[i].BiayaBulanan
		}
	}

	fmt.Println("\nPengeluaran per kategori:")
	for kategori, biaya := range kategoriMap {
		persentase := (biaya / total) * 100
		fmt.Printf("- %s: Rp %.2f (%.1f%%)\n", kategori, biaya, persentase)
	}
}

func getTanggalHariIni() Tanggal {
	now := time.Now()
	return Tanggal{
		Hari:  now.Day(),
		Bulan: int(now.Month()),
		Tahun: now.Year(),
	}
}

func hitungSelisihHari(tanggal1, tanggal2 Tanggal) int {
	// hitung selisih hari
	// Asumsi: tanggal dalam bulan dan tahun yang sama
	return tanggal2.Hari - tanggal1.Hari
}

func pengingatPembayaran() {
	if daftarLangganan.Count == 0 {
		fmt.Println("Tidak ada langganan yang tersimpan.")
		return
	}

	fmt.Println("\n=== PENGINGAT PEMBAYARAN ===")
	tanggalHariIni := getTanggalHariIni()

	fmt.Println("Langganan yang akan jatuh tempo dalam 7 hari:")
	adaPengingat := false

	for i := 0; i < daftarLangganan.Count; i++ {
		if daftarLangganan.Data[i].Status == "Aktif" {
			selisih := hitungSelisihHari(tanggalHariIni, daftarLangganan.Data[i].TanggalBayar)
			if selisih >= 0 && selisih <= 7 {
				fmt.Printf("- %s: Jatuh tempo %d/%d/%d (dalam %d hari) - Rp %.2f\n",
					daftarLangganan.Data[i].Nama,
					daftarLangganan.Data[i].TanggalBayar.Hari,
					daftarLangganan.Data[i].TanggalBayar.Bulan,
					daftarLangganan.Data[i].TanggalBayar.Tahun,
					selisih,
					daftarLangganan.Data[i].BiayaBulanan)
				adaPengingat = true
			}
		}
	}

	if !adaPengingat {
		fmt.Println("Tidak ada pembayaran yang akan jatuh tempo dalam 7 hari ke depan.")
	}
}

func rekomendasiPenghematan() {
	if daftarLangganan.Count == 0 {
		fmt.Println("Tidak ada langganan yang tersimpan.")
		return
	}

	fmt.Println("\n=== REKOMENDASI PENGHEMATAN ===")

	// sort by biaya tertinggi
	var tempData [MAX_SUBSCRIPTIONS]Langganan
	var tempCount int

	// copy data array sementara
	for i := 0; i < daftarLangganan.Count; i++ {
		if daftarLangganan.Data[i].Status == "Aktif" {
			tempData[tempCount] = daftarLangganan.Data[i]
			tempCount++
		}
	}

	// descending by biaya
	for i := 0; i < tempCount-1; i++ {
		maxIdx := i
		for j := i + 1; j < tempCount; j++ {
			if tempData[j].BiayaBulanan > tempData[maxIdx].BiayaBulanan {
				maxIdx = j
			}
		}
		if maxIdx != i {
			tempData[i], tempData[maxIdx] = tempData[maxIdx], tempData[i]
		}
	}

	total := hitungTotalPengeluaran()
	fmt.Printf("Total pengeluaran saat ini: Rp %.2f\n\n", total)

	fmt.Println("Rekomendasi langganan yang bisa dihentikan untuk menghemat biaya:")
	fmt.Println("(Diurutkan dari biaya tertinggi)")

	var penghematan float64
	for i := 0; i < tempCount; i++ {
		penghematan += tempData[i].BiayaBulanan
		fmt.Printf("%d. %s - Rp %.2f\n", i+1, tempData[i].Nama, tempData[i].BiayaBulanan)
		fmt.Printf("   Penghematan jika dihentikan: Rp %.2f (%.1f%% dari total)\n\n",
			tempData[i].BiayaBulanan, (tempData[i].BiayaBulanan/total)*100)
	}
}

func main() {

	daftarLangganan.Count = 0

	var pilihan int
	for {
		tampilkanMenu()
		pilihan = bacaInt("Pilih menu: ")

		switch pilihan {
		case 1:
			tambahLangganan()
		case 2:
			lihatSemuaLangganan()
		case 3:
			cariLangganan()
		case 4:
			editLangganan()
		case 5:
			hapusLangganan()
		case 6:
			urutkanLangganan()
		case 7:
			laporanKeuangan()
		case 8:
			pengingatPembayaran()
		case 9:
			rekomendasiPenghematan()
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}

		fmt.Println("\nTekan Enter untuk melanjutkan...")
		bufio.NewReader(os.Stdin).ReadLine()
	}
}
