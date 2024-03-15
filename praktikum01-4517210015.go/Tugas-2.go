package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mahasiswa struct {
	Nama    string
	NPM     string
	Jurusan string
}

func main() {
	mahasiswas := make(map[string]Mahasiswa)

	mahasiswas["4517210015"] = Mahasiswa{"Garin Hanggario", "4517210015", "Teknik Informatika"}
	mahasiswas["4520210015"] = Mahasiswa{"Paul McCartney", "4520210015", "Teknik Informatika"}

	fmt.Println("Daftar Mahasiswa:")
	for _, mahasiswa := range mahasiswas {
		fmt.Println(mahasiswa.Nama)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan NPM mahasiswa yang ingin dicari: ")
	npmInput, _ := reader.ReadString('\n')
	npm := strings.TrimSpace(npmInput)

	if mahasiswa, found := mahasiswas[npm]; found {
		fmt.Printf("Mahasiswa dengan NPM %s ditemukan:\n", npm)
		fmt.Println("Nama:", mahasiswa.Nama)
		fmt.Println("Jurusan:", mahasiswa.Jurusan)
	} else {
		fmt.Println("Mahasiswa dengan NPM", npm, "tidak ditemukan.")
	}
}
