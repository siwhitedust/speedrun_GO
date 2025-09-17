package main

import (
	"fmt"
	"time"
)

type Data struct {
	nama   string
	umur   int
	alamat string
}

func main() {
	fmt.Println("Masuk ke repo")
	fmt.Println("Masuk 2")
	DataOrang()
	con()
}

func DataOrang() {
	Budi := Data{
		nama:   "Budi",
		umur:   20,
		alamat: "Jl. Negara",
	}

	fmt.Println(Budi.nama)
}

func con() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	time.Sleep(2 * time.Second)
	go func() {
		fmt.Println("proses pengambilan data")
		data := <-ch
		fmt.Println("data yang diambil", data)
	}()
}
