package node

type Sepatu struct {
	Id      int
	Merk    string
	Kondisi string
	Ukuran  float32
}

type TableSepatu struct {
	Data Sepatu
	Next *TableSepatu
}
