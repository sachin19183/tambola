package ticketgen

import (
	"fmt"
	"math/rand"
	"time"
)

//GenerateTicket generates a tambola ticket
/*
1)each ticket has 3 rows & 9 columns.
2) genrate 5 random numbers between 1 and 9 to pick place for 1st row
3) Repeat the process for second row. This time eliminate the colum value already used in step 1
   If the no of colum left after elimination is less than 3, repeat step 3
4) For the third row, fill the five positions not filled yet
*/
func GenerateTicket() (map[int]int, map[int]int, map[int]int) {
	var colOccur [10]int
	var colRepl []int
	time.Sleep(time.Duration(850) * time.Millisecond)
	seconds := time.Now().UnixNano()
	//fmt.Println("Before init, seed value: ", seconds)
	rand.Seed(seconds)
	row1, keyCount1 := populateRow()
	//fmt.Println(" ", keyCount1)
	for key := range row1 {
		//fmt.Printf("row1[%d] ", key)
		colOccur[key] = 1
	}
	//fmt.Println(row1)
	//fmt.Println("Going to generate colum indexes for row2")
	row2, keyCount2 := populateRow()
	//fmt.Println(" ", keyCount2)
	for key := range row2 {
		//fmt.Printf("row2[%d] ", key)
		if colOccur[key] == 1 {
			colOccur[key] = 12
		} else {
			colOccur[key] = 2
		}
	}
	fmt.Println("")
	//fmt.Println("Going to generate colum indexes for row3")
	row3, keyCount3 := populateRow()
	//fmt.Println(" ", keyCount3)
	for key := range row3 {
		//fmt.Printf("row3[%d] ", key)
		if colOccur[key] == 12 {
			colRepl = append(colRepl, key)
			colOccur[key] = 123
		} else if colOccur[key] == 1 {
			colOccur[key] = 13
		} else if colOccur[key] == 2 {
			colOccur[key] = 23
		} else {
			colOccur[key] = 3
		}
	}

	// All the colum position are generated. now we cross check if any colum was not included in the three rows
	// There could be two anomalies:
	// 1. The total no of unique key for a row was less than 5 (due to numbers getting repeated during generation)
	// 2. a particular colum present in all three rows.
	var keyCoverage = map[int]bool{1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false, 9: false}
	for key := range keyCoverage {
		_, ok1 := row1[key]
		if ok1 == true {
			keyCoverage[key] = true
		}
		_, ok2 := row2[key]
		if ok2 == true {
			keyCoverage[key] = true
		}
		_, ok3 := row3[key]
		if ok3 == true {
			keyCoverage[key] = true
		}
	}
	// Now we check from col index 1 to 9 , which col is not assigned to any rows.
	// val = false for that key indicates this colm is yet to be assigned
	//
	i := 0
	lenColRepl := len(colRepl)
	for key, val := range keyCoverage {
		//fmt.Printf("\n keyCov[%d] = %v", key, val)
		if val == false {
			if keyCount1 < 5 || keyCount2 < 5 || keyCount3 < 5 {
				if keyCount1 < 5 {
					row1[key] = 0
					keyCount1++
				}
				if keyCount2 < 5 {
					row2[key] = 0
					keyCount2++
				}
				if keyCount3 < 5 {
					row3[key] = 0
					keyCount3++
				}
			} else if lenColRepl > 0 {
				tmpKey := colRepl[i]
				//fmt.Printf("\n Key present in all three rows is %d\n", tmpKey)
				delete(row3, tmpKey)
				row3[key] = 0
				i++
				lenColRepl--
			} else {
				fmt.Println("Do something extra")

			}
		}

	}
	// Even after utilizing all columns, there could be scenario that the col occupation count for a row is less than 5
	// check that and make the occupancy count to 5 here
	if keyCount1 < 5 {
		for i := 1; i < 10; i++ {
			numToGuess := rand.Intn(10)
			fmt.Print(numToGuess, " ")
			if _, ok := row1[numToGuess]; ok {
				continue
			} else if numToGuess == 0 {
				continue
			} else {
				row1[numToGuess] = 0
				keyCount1++
				if keyCount1 == 5 {
					break
				}
			}
		}
	}
	if keyCount2 < 5 {
		for i := 1; i < 10; i++ {
			numToGuess := rand.Intn(10)
			fmt.Print(numToGuess, " ")
			if _, ok := row2[numToGuess]; ok {
				continue
			} else if numToGuess == 0 {
				continue
			} else {
				row2[numToGuess] = 0
				keyCount2++
				if keyCount2 == 5 {
					break
				}
			}
		}
	}
	if keyCount3 < 5 {
		for i := 1; i < 10; i++ {
			numToGuess := rand.Intn(10)
			fmt.Print(numToGuess, " ")
			if _, ok := row3[numToGuess]; ok {
				continue
			} else if numToGuess == 0 {
				continue
			} else {
				row3[numToGuess] = 0
				keyCount3++
				if keyCount3 == 5 {
					break
				}
			}
		}
	}
	/*
		fmt.Println("\nfinal outcome ")
		fmt.Println(" ")
		for key := range row1 {
			fmt.Printf("row1[%d] ", key)
		}
		fmt.Println(" ")
		for key := range row2 {
			fmt.Printf("row2[%d] ", key)
		}
		fmt.Println(" ")
		for key := range row3 {
			fmt.Printf("row3[%d] ", key)
		}
		fmt.Println(" ")
	*/
	// there could also be a case when all the colums are covered across three rows ,
	//yet the total number of colums allocated to each row is less than 5
	// So we have two different conditions:
	//1. All colums should be covered across three rows
	// 2. Each row should have 5 allocated colums
	/*
	 First check if all colums are covered or not.
	 if not, assign those colmns in first go to a row whose 5 colmn allocation is not reached.
	 if that is not the case, find a colm value which exist in all three rows and replace it with this value in the last row
	 if that is not the case then find a colmn value which exist in two rows and replace it with this value in the first found row

	*/
	return row1, row2, row3
}

func populateRow() (map[int]int, int) {

	var row = make(map[int]int)
	var keyCount int

	for i := 1; i < 10; i++ {

		numToGuess := rand.Intn(10)
		fmt.Print(numToGuess, " ")
		_, ok := row[numToGuess]
		if ok == true {
			//numToGuess is already a key in row1
			continue
		} else if numToGuess == 0 {
			continue
		} else {
			row[numToGuess] = 0
			keyCount++
			if keyCount >= 5 {
				break
			}
		}
	}
	//fmt.Println(row)
	//fmt.Println("Positions generated for this row with keyCount: ", keyCount)
	return row, keyCount
}
