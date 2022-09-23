package main

import (
	"fmt"
	"math"
	"time"
)

func findPrimes() {
	start := time.Now()
	var array [100]int
	j := 0
	for i := 0; i < 100; i++ {
		if i == 2 {
			array[j] = i
			j++
		}
		if i == 3 {
			array[j] = i
			j++
		}
		if i == 5 {
			array[j] = i
			j++
		}
		if i == 7 {
			array[j] = i
			j++
		}
		if i == 11 {
			array[j] = i
			j++
		}
		if i == 13 {
			array[j] = i
			j++
		}
		if i == 17 {
			array[j] = i
			j++
		}
		if i == 19 {
			array[j] = i
			j++
		}
		if i == 23 {
			array[j] = i
			j++
		}
		if i == 29 {
			array[j] = i
			j++
		}
		if i == 31 {
			array[j] = i
			j++
		}
		if i == 37 {
			array[j] = i
			j++
		}
		if i == 41 {
			array[j] = i
			j++
		}
		if i == 43 {
			array[j] = i
			j++
		}
		if i == 47 {
			array[j] = i
			j++
		}
		if i == 53 {
			array[j] = i
			j++
		}
		if i == 59 {
			array[j] = i
			j++
		}
		if i == 61 {
			array[j] = i
			j++
		}
		if i == 67 {
			array[j] = i
			j++
		}
		if i == 71 {
			array[j] = i
			j++
		}
		if i == 73 {
			array[j] = i
			j++
		}
		if i == 79 {
			array[j] = i
			j++
		}
		if i == 83 {
			array[j] = i
			j++
		}
		if i == 89 {
			array[j] = i
			j++
		}
		if i == 97 {
			array[j] = i
			j++
		}

	}
	fmt.Println(time.Since(start))
	for i := 0; i < j; i++ {
		fmt.Println(array[i])
	}
}

func primeFind(value int) {
	start := time.Now()
	f := make([]bool, value)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if f[i] == false {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	duration := time.Since(start)
	for i := 2; i < value; i++ {
		if f[i] == false {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println("")
	fmt.Println(duration)
}
