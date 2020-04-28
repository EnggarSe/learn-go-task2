package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("===========================Dibawah Merupakan Tugas Bilangan Genap=====================================")
	BilanganGenap()
	fmt.Println("======================================================================================================")
	fmt.Println("===========================Dibawah Merupakan Tugas Bilangan Ganjil=====================================")
	BilanganGanjil()
	fmt.Println("======================================================================================================")
	fmt.Println("===========================Dibawah Merupakan Tugas Bilangan Prima=====================================")
	BilanganPrima()
	fmt.Println("======================================================================================================")
}
func BilanganGenap() {
	array1 := make([]int, 30)
	var countGenap int
	for i := 0; i < len(array1)-10; i++ {
		array1[i] = rand.Intn(30)
	}
	fmt.Println(array1)
	for j := 0; j < len(array1); j++ {

		if array1[j]%2 == 0 && array1[j] != 0 {
			fmt.Println("Bilangan Genap 	: ", array1[j])
			countGenap++
		}
	}
	fmt.Println("Jumlah Bilangan Genap", countGenap)
}

func BilanganGanjil() {
	array1 := make([]int, 30)
	var countGanjil int
	for i := 0; i < len(array1)-10; i++ {
		array1[i] = rand.Intn(30)
	}
	fmt.Println(array1)
	for j := 0; j < len(array1); j++ {

		if array1[j]%2 == 1 {
			fmt.Println("Bilangan Genap 	: ", array1[j])
			countGanjil++
		}
	}
	fmt.Println("Jumlah Bilangan Ganjil", countGanjil)
}

func BilanganPrima() {
	array1 := make([]int, 26)
	jumlah := 0
	for i := 0; i < len(array1); i++ {
		array1[i] = i
	}
	fmt.Println(array1)

	for j := 1; j < len(array1); j++ {
		count := 0
		for y := 1; y <= array1[j]; y++ {
			if array1[j]%y == 0 {
				count++
			}
		}
		if count == 2 && array1[j] > 1 {
			jumlah++
			fmt.Println("Bilangan Prima :", array1[j])

		}
	}
	fmt.Println("Jumlah Bilangan Prima Sebanyak", jumlah)
}
