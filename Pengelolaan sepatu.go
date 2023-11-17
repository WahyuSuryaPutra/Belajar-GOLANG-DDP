package main

import "fmt"

type ukuran struct {
	us int
	eu int
	id int
	jp int
}

type sepatu struct {
	merk   string
	harga  int
	stock  int
	margin int
	warna  []string
	ukuran ukuran
}

func (s sepatu) hargaJual() int {
	return s.harga + s.margin
}

func (s sepatu) totalDuit() int {
	return s.hargaJual() * s.stock
}

func main() {
	daftarUkuran := ukuran{
		us: 10,
		eu: 42,
		id: 42,
		jp: 23,
	}

	daftarSepatu := []sepatu{
		{
			merk:   "Nike",
			harga:  2500000,
			margin: 200000,
			stock:  5,
			warna:  []string{"merah", "indigo", "putih"},
			ukuran: daftarUkuran,
		},
		{
			merk:   "Adidas", 
			harga:  2000000,
			margin: 200000,
			stock:  7,
			warna:  []string{"merah", "indigo", "putih"},
			ukuran: daftarUkuran,
		},
	}

	for _, s := range daftarSepatu {
		fmt.Println("Merk:", s.merk)
		fmt.Println("Harga:", s.harga)
		fmt.Println("Ukuran US:", s.ukuran.us)
		fmt.Println("Ukuran EU:", s.ukuran.eu)
		fmt.Println("Ukuran ID:", s.ukuran.id)
		fmt.Println("Ukuran JP:", s.ukuran.jp)
		fmt.Println("Harga Jual:", s.hargaJual()) 
		fmt.Println("Total Uang:", s.totalDuit()) 
	}
}
