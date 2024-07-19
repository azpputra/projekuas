package controller

import (
	"sepatu/model"
	"sepatu/node"
)

func InsertSepatu(id int, merk string, kondisi string, ukuran float32) bool {
	if id > 0 && merk != "" {
		model.InsertSepatu(id, merk, kondisi, ukuran)
		return true
	} else {
		return false
	}
}

func UpdateSepatu(id int, merk string, kondisi string, ukuran float32) bool {
	if id > 0 && merk != "" {
		sepatu := node.Sepatu{
			Id:      id,
			Merk:    merk,
			Kondisi: kondisi,
			Ukuran:  ukuran,
		}
		return model.UpdateSepatu(sepatu)
	}
	return false
}

func SearchSepatu(id int, merk string, kondisi string, ukuran float32) *node.Sepatu {
	if id > 0 {
		addr := model.SearchSepatu(id)
		if addr != nil {
			return &addr.Data
		}
	}
	return nil
}

func DeleteSepatu(id int) bool {
	if id > 0 {
		return model.DeleteSepatu(id)
	}
	return false
}

func ReadAllSepatu() *node.TableSepatu {
	return model.ReadAllSepatu()
}
