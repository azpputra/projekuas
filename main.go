package main

import (
	"fmt"
	"sepatu/view"
)

func menu() {
	for {
		fmt.Println("Menu Program")
		fmt.Println("1. Insert Sepatu")
		fmt.Println("2. Update Sepatu")
		fmt.Println("3. Delete Sepatu")
		fmt.Println("4. Read All Sepatu")
		fmt.Println("5. Search Sepatu")
		fmt.Println("6. Exit")
		fmt.Println("---------------------")
		fmt.Print("masukkan menu pilihan anda: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			view.InsertSepatu()
		case 2:
			view.UpdateSepatu()
		case 3:
			view.DeleteSepatu()
		case 4:
			view.ReadAllSepatu()
		case 5:
			view.SearchSepatu()
		case 6:
			fmt.Println("exit program")
			return
		}
	}
}

func main() {
	menu()

	
}
