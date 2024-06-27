package main

import "fmt"

func main() {
	for panjangSegitiga := 1; panjangSegitiga <= 10; panjangSegitiga++ {
		for cetakSpasi := 1; cetakSpasi <= 10-panjangSegitiga; cetakSpasi++ {
			fmt.Print(" ")
		}
		for cetakBintang := 1; cetakBintang <= panjangSegitiga*2-1; cetakBintang++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
