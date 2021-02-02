package main

import (
	"fmt"

	ticketgen "github.com/tambola/ticketGen"
)

func main() {

	fmt.Println("Ticket 1:")
	row1Pos, row2Pos, row3Pos := ticketgen.GenerateTicket()
	fmt.Println("\nTicket 2:")
	row4Pos, row5Pos, row6Pos := ticketgen.GenerateTicket()
	fmt.Println("\nTicket 3:")
	row7Pos, row8Pos, row9Pos := ticketgen.GenerateTicket()
	fmt.Println("\nTicket 4:")
	row10Pos, row11Pos, row12Pos := ticketgen.GenerateTicket()
	fmt.Println("\nTicket 5:")
	row13Pos, row14Pos, row15Pos := ticketgen.GenerateTicket()
	fmt.Println("\nTicket 6:")
	row16Pos, row17Pos, row18Pos := ticketgen.GenerateTicket()

	arrPos := []map[int]int{row1Pos, row2Pos, row3Pos, row4Pos, row5Pos, row6Pos, row7Pos, row8Pos, row9Pos, row10Pos, row11Pos, row12Pos, row13Pos, row14Pos, row15Pos, row16Pos, row17Pos, row18Pos}

	//store occupency count of each colum
	var minCol []int //shall contain those columns whose occupency count is less than allowed value
	var majCol []int // shall contain those columns whose occupency count is more than allowed value
	var colCount [9]int
	allowedColCount := [9]int{9, 10, 10, 10, 10, 10, 10, 10, 11}
	for cl := 0; cl < 9; cl++ {
		for rw := 0; rw < 18; rw++ {
			tmpMap := arrPos[rw]
			_, ok := tmpMap[cl+1]
			if ok == true {
				colCount[cl]++
			}
		}
		//fmt.Printf(" column %d count is %d\n . ", cl, colCount[cl])
		if colCount[cl] < allowedColCount[cl] {
			minCol = append(minCol, cl)
			//fmt.Printf("%d column is in deficiency\n", cl)
		} else if colCount[cl] > allowedColCount[cl] {
			majCol = append(majCol, cl)
			//fmt.Printf("%d column is in excess \n", cl)
		} else {
			fmt.Printf("%d column is ok\n ", cl)
		}
	}
	fmt.Println("majcol is ", majCol)

	//	for cm := range majCol {
	//fmt.Printf(" %d ", cm)
	//	}
	fmt.Println("minCol is ", minCol)
	//	for cn := range minCol {
	//		fmt.Printf(" %d ", cn)
	//	}
	tkt1 := arrPos[:3]

	tkt2 := arrPos[3:6]
	tkt3 := arrPos[6:9]
	tkt4 := arrPos[9:12]
	tkt5 := arrPos[12:15]
	tkt6 := arrPos[15:18]

	/******************  Ticket 1 **********************/
	for iter1 := range majCol {
		count1 := 0
		c1 := majCol[iter1]
		fmt.Printf("Procesing %d in tkt1 \n", c1)
		for rw := 0; rw < 3; rw++ {
			tmpMap := tkt1[rw]
			_, ok := tmpMap[c1+1]
			if ok == true {
				count1++
			}
		}
		if count1 >= 2 && colCount[c1] > allowedColCount[c1] {
			//to maintain the row count as 5, add another key from minority col
			//breakPoint := false
			//fmt.Printf(" %d is present %d times in tikt1 \n", c1, colCount[c1])
			for iter2 := range minCol {
				count2 := 0
				c2 := minCol[iter2]
				for rw := 0; rw < 3; rw++ {
					tmpMap := tkt1[rw]
					_, ok2 := tmpMap[c2+1]
					if ok2 == true {
						count2++
					}
				}
				if count2 <= 2 && colCount[c2] < allowedColCount[c2] && colCount[c1] > allowedColCount[c1] {
					//find a row where this key is not present
					//fmt.Printf(" %d is presnt only once and its count is %d \n", c2, colCount[c2])
					for rw := 0; rw < 3; rw++ {
						tmpMap := tkt1[rw]
						_, ok2 := tmpMap[c2+1]
						_, ok1 := tmpMap[c1+1]
						if ok2 == false && ok1 == true { // excess col is present and min col is absent, replace.
							fmt.Printf("Going to  replace %d with %d in row %d \n", c1, c2, rw)
							delete(tmpMap, c1+1)
							colCount[c1]--
							tmpMap[c2+1] = 0
							colCount[c2]++

							break // only one key to be replaced in this for loop

						}
					}
				} else {
					fmt.Printf(" %d is present more than twice in this ticket \n", c2)
				}

			}
		}

	}

	/******************  Ticket 2 **********************/
	for iter1 := range majCol {
		count1 := 0
		c1 := majCol[iter1]
		fmt.Printf("Procesing %d in tkt2 \n", c1)
		for rw := 0; rw < 3; rw++ {
			tmpMap := tkt2[rw]
			_, ok := tmpMap[c1+1]
			if ok == true {
				count1++
			}
		}
		if count1 >= 2 && colCount[c1] > allowedColCount[c1] {
			//to maintain the row count as 5, add another key from minority col
			//breakPoint := false
			//fmt.Printf(" %d is present %d times in tikt2 \n", c1, colCount[c1])
			for iter2 := range minCol {
				count2 := 0
				c2 := minCol[iter2]
				for rw := 0; rw < 3; rw++ {
					tmpMap := tkt2[rw]
					_, ok2 := tmpMap[c2+1]
					if ok2 == true {
						count2++
					}
				}
				if count2 <= 2 && colCount[c2] < allowedColCount[c2] && colCount[c1] > allowedColCount[c1] {
					//find a row where this key is not present
					//fmt.Printf(" %d is presnt only once and its count is %d \n", c2, colCount[c2])
					for rw := 0; rw < 3; rw++ {
						tmpMap := tkt2[rw]
						_, ok2 := tmpMap[c2+1]
						_, ok1 := tmpMap[c1+1]
						if ok2 == false && ok1 == true {
							fmt.Printf("Going to  replace %d with %d in row %d \n", c1, c2, rw)
							delete(tmpMap, c1+1)
							colCount[c1]--
							tmpMap[c2+1] = 0
							colCount[c2]++

							break // only one key to be replaced in this for loop

						}
					}
				} else {
					fmt.Printf(" %d is present more than twice in this ticket \n", c2)
				}

			}
		}

	}
	/******************  Ticket 3 **********************/
	for iter1 := range majCol {
		count1 := 0
		c1 := majCol[iter1]
		fmt.Printf("Procesing %d in tkt3 \n", c1)
		for rw := 0; rw < 3; rw++ {
			tmpMap := tkt3[rw]
			_, ok := tmpMap[c1+1]
			if ok == true {
				count1++
			}
		}
		if count1 >= 2 && colCount[c1] > allowedColCount[c1] {
			//to maintain the row count as 5, add another key from minority col
			//breakPoint := false
			fmt.Printf(" %d is present three times in tikt3 \n", c1)
			for iter2 := range minCol {
				count2 := 0
				c2 := minCol[iter2]
				for rw := 0; rw < 3; rw++ {
					tmpMap := tkt3[rw]
					_, ok2 := tmpMap[c2+1]
					if ok2 == true {
						count2++
					}
				}
				if count2 <= 2 && colCount[c2] < allowedColCount[c2] && colCount[c1] > allowedColCount[c1] {
					//find a row where this key is not present
					//fmt.Printf(" %d is presnt only once and its count is %d \n", c2, colCount[c2])
					for rw := 0; rw < 3; rw++ {
						tmpMap := tkt3[rw]
						_, ok2 := tmpMap[c2+1]
						_, ok1 := tmpMap[c1+1]
						if ok2 == false && ok1 == true {
							fmt.Printf("Going to  replace %d with %d in row %d \n", c1, c2, rw)
							delete(tmpMap, c1+1)
							colCount[c1]--
							tmpMap[c2+1] = 0
							colCount[c2]++

							break // only one key to be replaced in this for loop

						}
					}
				} else {
					fmt.Printf(" %d is present more than twice in this ticket \n", c2)
				}

			}
		}
	}
	/******************  Ticket 4 **********************/
	for iter1 := range majCol {
		count1 := 0
		c1 := majCol[iter1]
		fmt.Printf("Procesing %d in tkt4 \n", c1)
		for rw := 0; rw < 3; rw++ {
			tmpMap := tkt4[rw]
			_, ok := tmpMap[c1+1]
			if ok == true {
				count1++
			}
		}
		if count1 >= 2 && colCount[c1] > allowedColCount[c1] {
			//to maintain the row count as 5, add another key from minority col
			//breakPoint := false
			fmt.Printf(" %d is present three times in tikt4 \n", c1)
			for iter2 := range minCol {
				count2 := 0
				c2 := minCol[iter2]
				for rw := 0; rw < 3; rw++ {
					tmpMap := tkt4[rw]
					_, ok2 := tmpMap[c2+1]
					if ok2 == true {
						count2++
					}
				}
				if count2 <= 2 && colCount[c2] < allowedColCount[c2] && colCount[c1] > allowedColCount[c1] {
					//find a row where this key is not present
					//fmt.Printf(" %d is presnt only once and its count is %d \n", c2, colCount[c2])
					for rw := 0; rw < 3; rw++ {
						tmpMap := tkt4[rw]
						_, ok2 := tmpMap[c2+1]
						_, ok1 := tmpMap[c1+1]
						if ok2 == false && ok1 == true {
							fmt.Printf("Going to  replace %d with %d in row %d \n", c1, c2, rw)
							delete(tmpMap, c1+1)
							colCount[c1]--
							tmpMap[c2+1] = 0
							colCount[c2]++

							break // only one key to be replaced in this for loop

						}
					}
				} else {
					fmt.Printf(" %d is present more than twce in this ticket \n", c2)
				}

			}
		}
	}

	/******************  Ticket 5 **********************/
	for iter1 := range majCol {
		count1 := 0
		c1 := majCol[iter1]
		fmt.Printf("Procesing %d in tkt5 \n", c1)
		for rw := 0; rw < 3; rw++ {
			tmpMap := tkt5[rw]
			_, ok := tmpMap[c1+1]
			if ok == true {
				count1++
			}
		}
		if count1 >= 2 && colCount[c1] > allowedColCount[c1] {
			//to maintain the row count as 5, add another key from minority col
			//breakPoint := false
			//fmt.Printf(" %d is present three times in tikt5 \n", c1)
			for iter2 := range minCol {
				count2 := 0
				c2 := minCol[iter2]
				for rw := 0; rw < 3; rw++ {
					tmpMap := tkt5[rw]
					_, ok2 := tmpMap[c2+1]
					if ok2 == true {
						count2++
					}
				}
				if count2 <= 2 && colCount[c2] < allowedColCount[c2] && colCount[c1] > allowedColCount[c1] {
					//find a row where this key is not present
					//		fmt.Printf(" %d is presnt only once and its count is %d \n", c2, colCount[c2])
					for rw := 0; rw < 3; rw++ {
						tmpMap := tkt5[rw]
						_, ok2 := tmpMap[c2+1]
						_, ok1 := tmpMap[c1+1]
						if ok2 == false && ok1 == true {
							fmt.Printf("Going to  replace %d with %d in row %d \n", c1, c2, rw)
							delete(tmpMap, c1+1)
							colCount[c1]--
							tmpMap[c2+1] = 0
							colCount[c2]++

							break // only one key to be replaced in this for loop

						}
					}
				} else {
					fmt.Printf(" %d is present more than twce in this ticket \n", c2)
				}

			}
		}
	}

	/******************  Ticket 6 **********************/
	for iter1 := range majCol {
		count1 := 0
		c1 := majCol[iter1]
		fmt.Printf("Procesing %d in tkt6 \n", c1)
		for rw := 0; rw < 3; rw++ {
			tmpMap := tkt6[rw]
			_, ok := tmpMap[c1+1]
			if ok == true {
				count1++
			}
		}
		if count1 >= 2 && colCount[c1] > allowedColCount[c1] {
			//to maintain the row count as 5, add another key from minority col
			//breakPoint := false
			fmt.Printf(" %d is present three times in tikt6 \n", c1)
			for iter2 := range minCol {
				count2 := 0
				c2 := minCol[iter2]
				for rw := 0; rw < 3; rw++ {
					tmpMap := tkt6[rw]
					_, ok2 := tmpMap[c2+1]
					if ok2 == true {
						count2++
					}
				}
				if count2 <= 2 && colCount[c2] < allowedColCount[c2] && colCount[c1] > allowedColCount[c1] {
					//find a row where this key is not present
					//	fmt.Printf(" %d is presnt only once and its count is %d \n", c2, colCount[c2])
					for rw := 0; rw < 3; rw++ {
						tmpMap := tkt6[rw]
						_, ok2 := tmpMap[c2+1]
						_, ok1 := tmpMap[c1+1]
						if ok2 == false && ok1 == true {
							fmt.Printf("Going to  replace %d with %d in row %d \n", c1, c2, rw)
							delete(tmpMap, c1+1)
							colCount[c1]--
							tmpMap[c2+1] = 0
							colCount[c2]++

							break // only one key to be replaced in this for loop

						}
					}
				} else {
					fmt.Printf(" %d is present more than twice in this ticket \n", c2)
				}

			}
		}
	}
	fmt.Println(arrPos)
	fmt.Println(colCount)
	var ticket [18][9]int
	/*
		for ii := 0; ii < 18; ii++ {
			for jj := 0; jj < 9; jj++ {
				ticket[ii][jj] = -1
			}
		}
	*/
	randSlice1 := []int{3, 5, 2, 1, 6, 8, 7, 9, 4}
	randSlice2 := []int{13, 15, 12, 14, 11, 16, 18, 17, 19, 10}
	randSlice3 := []int{23, 25, 24, 22, 20, 21, 26, 28, 27, 29}
	randSlice4 := []int{34, 30, 35, 32, 31, 36, 38, 37, 39, 33}
	randSlice5 := []int{43, 45, 42, 41, 46, 40, 48, 47, 49, 44}
	randSlice6 := []int{55, 53, 50, 52, 51, 56, 58, 54, 59, 57}
	randSlice7 := []int{63, 64, 65, 62, 61, 66, 68, 67, 69, 60}
	randSlice8 := []int{73, 75, 72, 74, 71, 70, 76, 78, 77, 79}
	randSlice9 := []int{81, 83, 85, 90, 82, 80, 84, 86, 88, 87, 89}
	randomArr := [9][]int{randSlice1, randSlice2, randSlice3, randSlice4, randSlice5, randSlice6, randSlice7, randSlice8, randSlice9}

	for iter := range randomArr {
		//fmt.Printf(" %d ", iter)
		//tmpArr := randomArr[iter]
		randTmp := randomArr[iter]
		fmt.Println(randTmp)
		iter2 := 0
		for rwc := 0; rwc < 18; rwc++ {
			tpMap := arrPos[rwc] // take the row snapshot
			//fmt.Println("row position snapshot", tpMap)
			_, ok := tpMap[iter+1] // check if the col position is marked in this row
			if ok == true {        // if yes, store the random number value for this col in this ticket position
				//fmt.Printf(" rw: %d| iter: %d | iter2: %d \n", rwc, iter, iter2)
				ticket[rwc][iter] = randTmp[iter2]
				tpMap[iter+1] = randTmp[iter2]
				iter2++
			}
		}
		//fmt.Printf("ticket after finishing colum %d ", iter)
		//fmt.Println(ticket)
		printArr(ticket)
	}

	// sort numbers in each ticket
	ticket1 := ticket[:3]
	sortTicket(ticket1)
	ticket2 := ticket[3:6]
	sortTicket(ticket2)
	ticket3 := ticket[:3]
	sortTicket(ticket3)
	ticket4 := ticket[3:6]
	sortTicket(ticket4)
	ticket5 := ticket[:3]
	sortTicket(ticket5)
	ticket6 := ticket[3:6]
	sortTicket(ticket6)

	fmt.Println("")
	printArr(ticket)
	fmt.Println(arrPos)
}
func printArr(ad [18][9]int) {
	for row := 0; row < 18; row++ {
		for col := 0; col < 9; col++ {
			fmt.Printf("%d ", ad[row][col])
		}
		fmt.Println(" ")
	}
	fmt.Println("--------------------------")
}

func sortTicket(ticket1 [][9]int) {
	fmt.Println(ticket1)
	fmt.Println("------------------------------------------------------------------")
	//fmt.Println(ticket)
	//fmt.Println("--------------------*********************")

	for col := 0; col < 9; col++ {
		if ticket1[0][col] != 0 {
			if ticket1[1][col] != 0 && ticket1[0][col] > ticket1[1][col] {
				tmp := ticket1[0][col]
				ticket1[0][col] = ticket1[1][col]
				ticket1[1][col] = tmp
			}
		}
		if ticket1[1][col] != 0 && ticket1[2][col] != 0 {
			if ticket1[1][col] > ticket1[2][col] {
				tmp := ticket1[1][col]
				ticket1[1][col] = ticket1[2][col]
				ticket1[2][col] = tmp
				if ticket1[0][col] != 0 {
					if ticket1[1][col] != 0 && ticket1[0][col] > ticket1[1][col] {
						tmp := ticket1[0][col]
						ticket1[0][col] = ticket1[1][col]
						ticket1[1][col] = tmp
					}
				}
			}
		}
		if ticket1[0][col] != 0 && ticket1[1][col] == 0 && ticket1[2][col] != 0 {
			if ticket1[0][col] > ticket1[2][col] {
				tmp := ticket1[0][col]
				ticket1[0][col] = ticket1[2][col]
				ticket1[2][col] = tmp
			}
		}
	}
	fmt.Println("Result is:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf(" %d ", ticket1[i][j])
		}
		fmt.Println(" ")
	}
	fmt.Println("------------------------------------------------------------------")
}
