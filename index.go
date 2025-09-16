package main

import "fmt"

type Data struct {
	nama   string
	umur   int
	alamat string
}

func main() {
	fmt.Println("Masuk ke repo")
	fmt.Println("Masuk 2")
	DataOrang()
}

func DataOrang() {
	Budi := Data{
		nama:   "Budi",
		umur:   20,
		alamat: "Jl. Negara",
	}

	fmt.Println(Budi)
}
