package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	breakLine        = "=============================================="
	alamatKantor     = "Office 8, lantai 15 SCBD Lot 28 Jl. Jendral Sudirman Kav. 52-53 Jaksel Kode POS 12190"
	backEndDeveloper = "Back End Developer"
	reason           = "Supaya menjadi mahir dalam bahasa Golang"
	errorMessage     = "Maaf terjadi kesalahan. Silahkan periksa kembali inputan kamu"
	dataNotFound     = "Data tidak ditemukan"
	title            = "ASSIGNMENT 1 GOLANG HACKTIVE8"
)

type Biodata struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	var bioData = []Biodata{
		{Absen: 1, Nama: "Zola Octanoviar", Alamat: alamatKantor, Pekerjaan: backEndDeveloper, Alasan: reason},
		{Absen: 2, Nama: "Dewo Shero", Alamat: alamatKantor, Pekerjaan: backEndDeveloper, Alasan: reason},
		{Absen: 3, Nama: "Ipan Taupik Rahman", Alamat: alamatKantor, Pekerjaan: backEndDeveloper, Alasan: reason},
		{Absen: 4, Nama: "Jeffri Alfred", Alamat: alamatKantor, Pekerjaan: backEndDeveloper, Alasan: reason},
		{Absen: 5, Nama: "Ahmad Dikri Alpakih", Alamat: alamatKantor, Pekerjaan: backEndDeveloper, Alasan: reason},
		{Absen: 6, Nama: "Thomas Andrianto", Alamat: alamatKantor, Pekerjaan: backEndDeveloper, Alasan: reason},
		{Absen: 7, Nama: "Arif", Alamat: alamatKantor, Pekerjaan: backEndDeveloper, Alasan: reason},
		{Absen: 8, Nama: "ZamZam", Alamat: alamatKantor, Pekerjaan: backEndDeveloper, Alasan: reason},
		{Absen: 9, Nama: "Ricky Susanto", Alamat: alamatKantor, Pekerjaan: backEndDeveloper, Alasan: reason},
	}
	fmt.Println(title)
	fmt.Println(breakLine)
	allArgs := os.Args[1:]
	if len(os.Args) > 1 {
		for i := 0; i < len(allArgs); i++ {
			absen, err := strconv.Atoi(allArgs[i])
			if err != nil {
				fmt.Println(errorMessage)
			} else {
				search := 0
				for i := 0; i < len(bioData); i++ {
					if bioData[i].Absen == absen {
						ShowData(bioData[i])
						search++
						break
					}
				}
				if search == 0 {
					fmt.Println(dataNotFound)
					fmt.Println(breakLine)
				}
			}
		}
	}
}
func ShowData(param Biodata) {
	fmt.Println("Nama :", param.Nama)
	fmt.Println("Alamat :", param.Alamat)
	fmt.Println("Pekerjaan :", param.Pekerjaan)
	fmt.Println("Alasan :", param.Alasan)
	fmt.Println(breakLine)
}
