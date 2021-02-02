package player

import "fmt"

// Player is a struct which stores relevant information about a tambola player
type Player struct {
	name        string
	firstRow    []int
	secRow      []int
	thirdRow    []int
	corner      []int
	isCorner    bool
	isFirstRow  bool
	isSecRow    bool
	isThirdRow  bool
	isFirstFive bool
	tkt         *[3][9]int
}

func initPlayer(name string) *Player {
	var p Player
	p.name = name
	var ticket = [3][9]int{{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{21, 22, 23, 24, 25, 26, 27, 28, 29},
		{31, 32, 33, 34, 35, 36, 37, 38, 39}}
	p.tkt = &ticket
	return &p
}

func (p Player) populateRow(rowNum int, row []int) {
	switch rowNum {
	case 1:
		for i := 0; i < 5; i++ {
			p.firstRow = append(p.firstRow, row[i])
		}
		p.corner = append(p.corner, row[0])
		p.corner = append(p.corner, row[4])

	case 2:
		for i := 0; i < 5; i++ {
			p.secRow = append(p.secRow, row[i])
		}

	case 3:
		for i := 0; i < 5; i++ {
			p.thirdRow = append(p.thirdRow, row[i])
		}
		p.corner = append(p.corner, row[0])
		p.corner = append(p.corner, row[4])

	default:
		fmt.Println("You should not be here")
	}

}

func (p Player) validateRowDone(rowNum int, numSeries map[int]bool) bool {
	validated := true
	switch rowNum {
	case 1:
		for i := 0; i < 5; i++ {
			if numSeries[p.firstRow[i]] == false {
				validated = false
				break
			}
		}
		return validated

	case 2:
		for i := 0; i < 5; i++ {
			if numSeries[p.secRow[i]] == false {
				validated = false
				break
			}
		}
		return validated

	case 3:
		for i := 0; i < 5; i++ {
			if numSeries[p.thirdRow[i]] == false {
				validated = false
				break
			}
		}
		return validated
	default:
		fmt.Println("Invalid value of row number.Returning false")
		return false
	}
}

func (p Player) validateCorner(numSeries map[int]bool) bool {
	validated := true
	for i := 0; i < 5; i++ {
		if numSeries[p.corner[i]] == false {
			validated = false
			break
		}
	}
	return validated
}
