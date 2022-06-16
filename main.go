package main

import "fmt"

var (
	aple    = 5.99
	pear    = 7
	myMoney = 23
)

func main() {

	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Println((9*float64(aple))+(8*float64(pear)), "грн")

	fmt.Println("2. Скільки груш ми можемо купити?")
	fmt.Println(int(float64(myMoney)/float64(pear)), "шт")

	fmt.Println("3. Скільки яблук ми можемо купити?")
	fmt.Println(int(float64(myMoney)/float64(aple)), "шт")

	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?")
	fmt.Println(2*(float64(aple)+float64(pear)) <= float64(myMoney))

}

// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
// 2. Скільки груш ми можемо купити?
// 3. Скільки яблук ми можемо купити?
// 4. Чи ми можемо купити 2 груші та 2 яблука?
