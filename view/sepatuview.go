package view

import (
	"fmt"
	"sepatu/controller"
)

func InsertSepatu() {
	var id int
	var merk, kondisi string
	var ukuran float32
	fmt.Print("Masukkan id sepatu : ")
	fmt.Scan(&id)
	fmt.Print("Masukkan merk sepatu : ")
	fmt.Scan(&merk)
	fmt.Print("Masukkan kondisi sepatu : ")
	fmt.Scan(&kondisi)
	fmt.Print("Masukkan ukuran sepatu : ")
	fmt.Scan(&ukuran)
	cek := controller.InsertSepatu(id, merk, kondisi, ukuran)
	if cek == true {
		fmt.Println("data berhasil diinput")
	} else {
		fmt.Println("data gagal diinput")
	}
}

func UpdateSepatu() {
	var id int
	var merk, kondisi string
	var ukuran float32
	fmt.Print("Masukkan id sepatu yang ingin diupdate: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan merk baru sepatu: ")
	fmt.Scan(&merk)
	fmt.Print("Masukkan kondisi baru sepatu: ")
	fmt.Scan(&kondisi)
	fmt.Print("Masukkan ukuran baru sepatu: ")
	fmt.Scan(&ukuran)
	cek := controller.UpdateSepatu(id, merk, kondisi, ukuran)
	if cek == true {
		fmt.Println("Data berhasil diupdate")
	} else {
		fmt.Println("Data gagal diupdate atau sepatu tidak ditemukan")
	}
}

func ReadAllSepatu() {
	sepatus := controller.ReadAllSepatu()
	if sepatus != nil {
		for sepatus.Next != nil {
			fmt.Println("Id Sepatu : ", sepatus.Next.Data.Id)
			fmt.Println("Merk Sepatu : ", sepatus.Next.Data.Merk)
			fmt.Println("Ukuran Sepatu : ", sepatus.Next.Data.Ukuran)
			fmt.Println("Kondisi Sepatu : ", sepatus.Next.Data.Kondisi)
			fmt.Println("---------------------")
			sepatus = sepatus.Next
		}
	}
}

func SearchSepatu() {
	var id int
	var merk, kondisi string
	var ukuran float32
	fmt.Print("Masukkan id sepatu yang ingin dicari: ")
	fmt.Scan(&id)
	sepatu := controller.SearchSepatu(id, merk, kondisi, ukuran)
	if sepatu != nil {
		fmt.Println("Id Sepatu : ", sepatu.Id)
		fmt.Println("Merk Sepatu : ", sepatu.Merk)
		fmt.Println("Ukuran Sepatu : ", sepatu.Ukuran)
		fmt.Println("Kondisi Sepatu : ", sepatu.Kondisi)
	} else {
		fmt.Println("Sepatu tidak ditemukan")
	}
}

func DeleteSepatu() {
	var id int
	fmt.Println("Masukkan id sepatu yang ingin dihapus: ")
	fmt.Scan(&id)
	cek := controller.DeleteSepatu(id)
	if cek == true {
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Data gagal dihapus atau sepatu tidak ditemukan")
	}
}
