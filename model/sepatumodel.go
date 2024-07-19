package model

import (
	"sepatu/database"
	"sepatu/node"
)

func InsertSepatu(id int, merk string, kondisi string, ukuran float32) {
	newDataSepatu := node.TableSepatu{
		Data: node.Sepatu{id, merk, kondisi, ukuran},
		Next: nil,
	}
	var tempLL *node.TableSepatu
	tempLL = &database.DBsepatu
	if database.DBsepatu.Next == nil {
		database.DBsepatu.Next = &newDataSepatu
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &newDataSepatu
	}
}

func ReadAllSepatu() *node.TableSepatu {
	var tempLL *node.TableSepatu
	tempLL = &database.DBsepatu
	if database.DBsepatu.Next == nil {
		return nil
	} else {
		return tempLL
	}
}

func DeleteSepatu(id int) bool {
	var tempLL *node.TableSepatu
	tempLL = &database.DBsepatu
	if tempLL.Next == nil {
		return false
	} else {
		for tempLL.Next != nil {
			if tempLL.Next.Data.Id == id {
				tempLL.Next = tempLL.Next.Next
				return true
			}
			tempLL = tempLL.Next
		}
		return false
	}

}

func SearchSepatu(id int) *node.TableSepatu {
	var tempLL *node.TableSepatu
	tempLL = database.DBsepatu.Next
	cek := false
	if database.DBsepatu.Next == nil {
		return nil
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				return tempLL
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			return nil
		}
	}
	return nil
}

func UpdateSepatu(sepatu node.Sepatu) bool {
	addr := SearchSepatu(sepatu.Id)
	if addr == nil {
		return false
	} else {
		addr.Data.Merk = sepatu.Merk
		addr.Data.Kondisi = sepatu.Kondisi
		addr.Data.Ukuran = sepatu.Ukuran
		return true
	}
}
