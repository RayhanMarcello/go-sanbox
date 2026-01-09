package main

type ProductJualan struct {
	Name  string
	Price int
}

func newProduct() *ProductJualan {
	return &ProductJualan{
		Name:  "cireng",
		Price: 10000,
	}
}

func ApplyDiscount(p *ProductJualan, persentage float32) float32 {
	afterDisc := float32(p.Price) * 0.1
	return afterDisc
}
