package main

import (
	"fmt"
	"math/rand"
)

var box [3][3]string

func printTable() {
	fmt.Println(".-.-.-.")
	for i:=0;i<3;i++ {
		for j:=0;j<3;j++ {
			if j == 0 {
				fmt.Printf("|")
			}
			text := box[i][j]
			if text == ""{
				text = " "
			}
			fmt.Printf("%s|", text)
		}
		fmt.Println()
		fmt.Println(".-.-.-.")
	}
}

func assignOpponentTurn() bool{
	for true {
		x := rand.Intn(3)
		y := rand.Intn(3)
		if box[y][x] == "" {
			box[y][x] = "X"
	        printTable()
			return calculateWin(x,y, "X")
		}
	}
	return false
}

func calculateWin(x int, y int, s string) bool {
var cnt int
// check vertical win
	for i:=0;i<3;i++ {
		if box[i][x] == s {
			cnt++
		}
	}
	if cnt == 3 {
		return true
	}
	
	// check horizontal win
	cnt = 0
	for i:=0;i<3;i++ {
		if box[y][i] == s {
			cnt++
		}
	}
	if cnt == 3 {
		return true
	}
	
	// check diagonal win (\)
	cnt = 0
	for i:=0;i<3;i++ {
		if box[i][i] == s {
			cnt++
		}
	}
	if cnt == 3 {
		return true
	}
	
	// check diagonal win (/)
	cnt = 0
	for i:=0;i<3;i++ {
		if box[i][2-i] == s {
			cnt++
		}
	}
	if cnt == 3 {
		return true
	}
	return false
}

func main() {
    var turn int
	printTable()
for true {
	var x, y int
	for true {
	for true {
		fmt.Println("input x")

         	_, err := fmt.Scanf("%d", &x)

         	if err != nil {
            		fmt.Println("x must be a number")
		} else if x < 1 || x > 4 {
			fmt.Println("x must be between 1 and 3")
		} else {
		    x--
			break
		}
	}
	for true {
		fmt.Println("input y")

         	_, err := fmt.Scanf("%d", &y)

         	if err != nil {
            		fmt.Println("y must be a number")
		} else if y < 1 || y > 4 {
			fmt.Println("y must be between 1 and 3")
		} else {
		    y--
			break
		}
	}
	if box[y][x] != "" {
		fmt.Println("that spot is already taken")
	} else {
		box[y][x] = "O"
		break
	}
	}
	printTable()
	win := calculateWin(x, y, "O")
	if win {
		fmt.Println("you win")
		break
	}
	turn++
	if turn == 9 {
	    fmt.Println("draw")
	    break
	}
	
	fmt.Println("Opponent's turn")
	win = assignOpponentTurn()
	if win {
		fmt.Println("you lose")
		break
	}
	turn++
}
}
