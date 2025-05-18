package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Assessment struct {
	ID int
	Nama string
	Date time.Time
	Scores []int
	TotalScore int
	Catatan string
}

var assessments []Assessment
var nextID int = 1
var reader = bufio.NewReader(os.Stdin)

var pertanyaan = []string{
	"Saya merasa puas dengan hidup saya : ",
	"Saya mampu mengatasi kesulitan dalam hidup saya : ",
	"Saya merasa termotivasi untuk melakukan aktivitas sehari-hari : ",
	"Saya bisa menikmati waktu bersama orang lain : ",
	"Saya merasa percaya diri dengan kemampuan saya : ",
	"Saya bisa tidur dengan nyenyak di malam hari : ",
	"Saya merasa bersyukur dengan apa yang saya miliki : ",
	"Saya merasa memiliki tujuan hidup yang jelas : ",
	"Saya sering merasa cemas tanpa alasan yang jelas : ",
	"Saya mengalami kesulitan untuk berkonsentrasi : ",
	"Saya merasa tidak berguna atau tidak berdaya : ",
	"Saya mudah tersinggung atau marah : ",
	"Saya merasa kelelahan bahkan tanpa aktivitas berat : ",
	"Saya merasa tidak tertarik pada hal-hal yang biasanya saya sukai : ",
	"Saya merasa hidup saya tidak berarti : ",
}

func main(){
	var pilihan int
	for {
		fmt.Println("\n=== Aplikasi Manajemen Kesehatan Mental ===")
		fmt.Println("1. Isi Self-Assessment DONE")
		fmt.Println("2. Kelola Data Assessment DONE")
		fmt.Println("3. Lihat Hasil dan Rekomendasi DONE")
		fmt.Println("4. Mencari Hasil Assessment berdasarkan ID DONE")
		fmt.Println("5. Urutkan Hasil Assessment")
		fmt.Println("6. Cetak Laporan DONE")
		fmt.Println("7. Tambah Catatan DONE")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)
		fmt.Println("===========================================")

		switch pilihan {
		case 1:
			isiAssessment()
		case 2:
			kelolaData()
		case 3:
			lihatHasil()
		case 4:
			cariAssessment()
		case 5:
			urutkanAssessment()
		case 6:
			cetakLaporan()
		case 7:
			tambahCatatan()
		case 8:
			fmt.Println("Sampai jumpa pada Self-Assessment berikutnya!")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func isiAssessment(){
	for{
		var ulang string

		fmt.Print("Masukkan nama Anda: ")
		nama, _ := reader.ReadString('\n')
		nama = strings.TrimSpace(nama)

		scores := make([]int, len(pertanyaan))
		total := 0

		fmt.Println("Jawab pertanyaan dibawah dengan skala 1 sampai 5\n 1 (Tidak pernah)\n 2 (Jarang)\n 3 (Netral)\n 4 (Sering)\n 5 (Selalu)")
		for i := 0; i < len(pertanyaan); i++ {
			fmt.Printf("%d. %s", i+1, pertanyaan[i])
			var nilai int
			for {
				fmt.Scanln(&nilai)
				if nilai >= 1 && nilai <= 5 {
					if i >= 8 {
						nilai = 6 - nilai // balik skor
					}
					scores[i] = nilai
					total += nilai
					break
				} else {
					fmt.Print("Masukkan angka antara 1 sampai 5 : ")
				}
			}
		}

		newAssessment := Assessment{
			ID: nextID,
			Nama: nama,
			Date: time.Now(),
			Scores: scores,
			TotalScore: total,

		}
		assessments = append(assessments, newAssessment)
		fmt.Println("Self-Assessment berhasil dengan ID : ", nextID)
		fmt.Println()
		nextID++
		fmt.Print("apakah ingin mengisi assessment lagi ? (y/n): ")
		fmt.Scanln(&ulang)
		if strings.ToLower(ulang) != "y" {
			break
		}
	}
}

func kelolaData(){
	for {
		var pilihan int
		fmt.Println("\n======= Kelola Data =======")
		fmt.Println("1. Hapus Data")
		fmt.Println("2. Edit Data")
		fmt.Println("3. Kembali")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		if pilihan == 3 {
			return
		}

		switch pilihan {
		case 1:
			hapusData()
		case 2:
			editData()
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func hapusData(){
	var id int
	fmt.Print("Masukkan ID yang ingin dihapus : ")
	fmt.Scanln(&id)

	index := cariIndexByID(id)
	if index == -1 {
		fmt.Println("ID tidak ditemukan")
		return
	}
	assessments = append(assessments[:index], assessments[index+1:]...)
	fmt.Println("Data berhasil dihapus")
}

func editData() {
	var id int
	fmt.Print("Masukkan ID yang ingin diedit: ")
	fmt.Scanln(&id)

	index := cariIndexByID(id)
	if index == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}

	fmt.Println("Nama lama:", assessments[index].Nama)
	fmt.Print("Masukkan nama baru (kosongkan jika tidak ingin mengubah): ")
	namaBaru, _ := reader.ReadString('\n')
	namaBaru = strings.TrimSpace(namaBaru)
	if namaBaru != "" {
		assessments[index].Nama = namaBaru
	}

	fmt.Println("Catatan lama:", assessments[index].Catatan)
	fmt.Print("Masukkan catatan baru (kosongkan jika tidak ingin mengubah): ")
	catatanBaru, _ := reader.ReadString('\n')
	catatanBaru = strings.TrimSpace(catatanBaru)
	if catatanBaru != "" {
		assessments[index].Catatan = catatanBaru
	}

	fmt.Println("Data berhasil diperbarui.")
}

func lihatHasil(){
	var nama string
	fmt.Print("\nMasukkan nama Anda: ")
	namaInput, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(namaInput)

	ditemukan := false
	for _, a := range assessments {
		if strings.EqualFold(a.Nama, nama) {
			tampilkanAssessment(a)
			rekomendasi(a.TotalScore)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ditemukan assessment dengan nama tersebut")
	}
}

func tampilkanAssessment(a Assessment) {
	fmt.Println("\n================ Hasil Assessment ================")
	fmt.Printf("%-12s: %d\n", "ID", a.ID)
	fmt.Printf("%-12s: %s\n", "Nama", a.Nama)
	fmt.Printf("%-12s: %s\n", "Tanggal", a.Date.Format("2006-01-02"))
	fmt.Printf("%-12s: %d\n", "Skor Total", a.TotalScore)
	fmt.Printf("%-12s: %s\n", "Catatan", a.Catatan)
	fmt.Println("\n--------------------------------------------------")
	fmt.Println()
}

func rekomendasi(score int){
	fmt.Println("📊 Rekomendasi untuk Anda:")
	switch {
	case score <= 30:
		fmt.Println("⚠️ KONDISI MENTAL ANDA CUKUP MENGKHAWATIRKAN ⚠️")
		fmt.Println("Kami sangat menyarankan Anda untuk segera berbicara dengan profesional seperti psikolog atau konselor.")
		fmt.Println("Anda tidak sendirian, selalu ada bantuan yang tersedia.")
	case score <= 50:
		fmt.Println("🌥️ KONDISI MENTAL ANDA CUKUP BAIK, NAMUN PERLU PENINGKATAN 🌥️")
		fmt.Println("Cobalah luangkan waktu untuk diri sendiri, bicarakan dengan orang terdekat, atau lakukan aktivitas yang menenangkan.")
		fmt.Println("Perhatikan pola tidur, makan, dan kebiasaan harian Anda.")
	default:
		fmt.Println("💖 KONDISI MENTAL ANDA SANGAT BAIK 💖")
		fmt.Println("Pertahankan kebiasaan positif dan terus jaga keseimbangan hidup.")
		fmt.Println("Jangan ragu untuk berbagi energi positif Anda ke orang sekitar!")
	}
	fmt.Println("====================================================")
}

func cariAssessment(){
	for {
		var metode, id int

		fmt.Println("\n======== Cari Assessment ========")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search")
		fmt.Println("3. Kembali")
		fmt.Println("==================================")
		fmt.Print("Pilih Metode : ")
		fmt.Scanln(&metode)
		if metode == 3 {
			return
		}
		fmt.Print("Masukkan ID yang dicari : ")
		fmt.Scanln(&id)

		switch metode {
		case 1:
			index := sequentialSearch(id)
			if index != -1 {
				tampilkanAssessment(assessments[index])
			} else {
				fmt.Println()
				fmt.Println("Assessment tidak ditemukan")
			}
		case 2:
			urutkanAssessmentsByID()
			index := binarySearch(id)
			if index != -1 {
				tampilkanAssessment(assessments[index])
			} else {
				fmt.Println("Assessment tidak ditemukan")
			}
		default:
			fmt.Println("Plihan tidak valid")
			continue
		}
	}
}

func sequentialSearch(id int) int { 
	for i := 0; i < len(assessments); i++ {
		if assessments[i].ID == id {
			return i
		}
	}
	return -1
}

func urutkanAssessmentsByID(){
	n := len(assessments)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if assessments[j].ID > assessments[j+1].ID {
				assessments[j], assessments[j+1] = assessments[j+1], assessments[j]
			}
		}
	}
}

func binarySearch(id int) int {
	var low int = 0
	var high int = len(assessments)-1
	for low <= high {
		mid := (low + high) / 2
		if assessments[mid].ID == id{
			return mid
		} else if assessments[mid].ID < id{
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func urutkanAssessment(){
	for {
		var metode int
		fmt.Println("\n======= Urutkan Assessment =======")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Println("3. Kembali")
		fmt.Print("Pilih metode: ")
		fmt.Scanln(&metode)

		if metode == 3 {
			return
		}

		var berdasarkan int
		fmt.Println("\nUrutkan berdasarkan:")
		fmt.Println("1. Skor Total")
		fmt.Println("2. Tanggal")
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&berdasarkan)

		var arah int
		fmt.Println("\nUrutan:")
		fmt.Println("1. Ascending (naik)")
		fmt.Println("2. Descending (turun)")
		fmt.Print("Pilih arah: ")
		fmt.Scanln(&arah)

		switch metode {
		case 1:
			selectionSort(berdasarkan, arah)
		case 2:
			insertionSort(berdasarkan, arah)
		default:
			fmt.Println("Metode tidak valid")
		}
	}
}

func insertionSort(berdasarkan int, arah int) {
	for i := 1; i < len(assessments); i++ {
		kunci := assessments[i]
		j := i - 1

		for j >= 0 {
			ganti := false
			switch berdasarkan {
			case 1: // skor
				if arah == 1 && assessments[j].TotalScore > kunci.TotalScore {
					ganti = true
				}
				if arah == 2 && assessments[j].TotalScore < kunci.TotalScore {
					ganti = true
				}
			case 2: // tgl
				if arah == 1 && assessments[j].Date.After(kunci.Date) {
					ganti = true
				}
				if arah == 2 && assessments[j].Date.Before(kunci.Date) {
					ganti = true
				}
			}
			if ganti {
				assessments[j+1] = assessments[j]
				j--
			} else {
				break
			}
		}
		assessments[j+1] = kunci
	}

	fmt.Println("\nData diurutkan:")
	for _, a := range assessments {
		tampilkanAssessment(a)
	}
}


func selectionSort(berdasarkan int, arah int) {
	n := len(assessments)
	for i := 0; i < n; i++ {
		index := i
		for j := i + 1; j < n; j++ {
			switch berdasarkan {
			case 1: // skor
				if arah == 1 && assessments[j].TotalScore < assessments[index].TotalScore {
					index = j
				}
				if arah == 2 && assessments[j].TotalScore > assessments[index].TotalScore {
					index = j
				}
			case 2: // tgl
				if arah == 1 && assessments[j].Date.Before(assessments[index].Date) {
					index = j
				}
				if arah == 2 && assessments[j].Date.After(assessments[index].Date) {
					index = j
				}
			}
		}
		assessments[i], assessments[index] = assessments[index], assessments[i]
	}

	fmt.Println("\nData diurutkan:")
	for _, a := range assessments {
		tampilkanAssessment(a)
	}
}

func cetakLaporan(){
	fmt.Println("\n======= Cetak Laporan =======")
	fmt.Print("Masukkan nama: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	var hasil []Assessment
	for _, a := range assessments {
		if strings.EqualFold(a.Nama, nama) {
			hasil = append(hasil, a)
		}
	}

	if len(hasil) == 0 {
		fmt.Println("Tidak ada data assessment dengan nama tersebut")
		return
	}

	fmt.Printf("\nMenampilkan %d hasil assessment untuk %s:\n", len(hasil), nama)
	for _, a := range hasil {
		tampilkanAssessment(a)
		rekomendasi(a.TotalScore)
	}

	// rata rata 30hr terahir
	now := time.Now()
	var total, count int
	for _, a := range hasil {
		if now.Sub(a.Date).Hours() <= 30*24 {
			total += a.TotalScore
			count++
		}
	}
	if count > 0 {
		fmt.Printf("\nRata-rata skor 30 hari terakhir untuk %s: %.2f\n", nama, float64(total)/float64(count))
		fmt.Println("====================================================")
	} else {
		fmt.Println("\nTidak ada assessment dalam 30 hari terakhir untuk nama ini")
	}
}

func tambahCatatan(){
	var id int
	fmt.Print("Masukkan ID assessment untuk menambahkan catatan: ")
	fmt.Scanln(&id)

	index := cariIndexByID(id)
	if index == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan catatan: ")
	note, _ := reader.ReadString('\n')
	note = strings.TrimSpace(note)
	assessments[index].Catatan = note
	fmt.Println("Catatan berhasil ditambahkan.")
}

// pencari ID
func cariIndexByID(id int) int {
    for i := 0; i < len(assessments); i++ {
        if assessments[i].ID == id {
            return i  
        }
    }
    return -1 
}
