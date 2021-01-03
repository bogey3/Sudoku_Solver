package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func longShot(board [][]int)bool{
	for y:=0; y<9; y++{
		for x:=0; x<9; x++{
			if board[y][x] == 0{
				possibles := findPossibles(board, y, x)
				for _, pos := range possibles{
					newBoard := make([][]int, len(board))
					for i := range board {
						newBoard[i] = make([]int, len(board[i]))
						copy(newBoard[i], board[i])
					}
					newBoard[y][x] = pos
					if properSolveBoard(newBoard){
						for i := range newBoard {
							board[i] = make([]int, len(newBoard[i]))
							copy(board[i], newBoard[i])
						}
						return true
					}
				}
			}
		}
	}
	return false
}


func intContains(arr []int, find int)bool{
	for _, v := range arr{
		if v == find{
			return true
		}
	}
	return false
}

func findIndex(arr []int, find int)(int, bool){
	for i, v := range arr{
		if v == find{
			return i, true
		}
	}
	return 0, false
}

func simpleDeduce(board [][]int, y int, x int)int{
	pool := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i:=1;i<10;i++{
		virt := []int{board[0][x], board[1][x], board[2][x], board[3][x], board[4][x], board[5][x], board[6][x], board[7][x], board[8][x]}
		box := []int{board[(y/3)*3 + 0][(x/3)*3 + 0], board[(y/3)*3 + 1][(x/3)*3 + 0], board[(y/3)*3 + 2][(x/3)*3 + 0], board[(y/3)*3 + 0][(x/3)*3 + 1], board[(y/3)*3 + 1][(x/3)*3 + 1], board[(y/3)*3 + 2][(x/3)*3 + 1], board[(y/3)*3 + 0][(x/3)*3 + 2], board[(y/3)*3 + 1][(x/3)*3 + 2], board[(y/3)*3 + 2][(x/3)*3 + 2]}
		total := countNumbers(board[y], i) + countNumbers(virt, i) + countNumbers(box, i)
		if total != 0{
			if loc, ok := findIndex(pool, i); ok {
				pool[loc] = pool[len(pool)-1]
				pool = pool[:len(pool)-1]
			}
		}
	}
	if len(pool) == 1 {
		return pool[0]
	}

	return 0
}


func complexDeduce(board [][]int, y int, x int)int{
	virt := [][]int{[]int{0, x}, []int{1, x}, []int{2, x}, []int{3, x}, []int{4, x}, []int{5, x}, []int{6, x}, []int{7, x}, []int{8, x}}
	horiz := [][]int{[]int{y, 0}, []int{y, 1}, []int{y, 2}, []int{y, 3}, []int{y, 4}, []int{y, 5}, []int{y, 6}, []int{y, 7}, []int{y, 8}}
	box := [][]int{[]int{(y/3)*3 + 0, (x/3)*3 + 0}, []int{(y/3)*3 + 1, (x/3)*3 + 0}, []int{(y/3)*3 + 2, (x/3)*3 + 0}, []int{(y/3)*3 + 0, (x/3)*3 + 1}, []int{(y/3)*3 + 1, (x/3)*3 + 1}, []int{(y/3)*3 + 2, (x/3)*3 + 1}, []int{(y/3)*3 + 0, (x/3)*3 + 2}, []int{(y/3)*3 + 1, (x/3)*3 + 2}, []int{(y/3)*3 + 2, (x/3)*3 + 2}}

	for _, set := range [][][]int{virt, horiz, box} {
		for num := 1; num < 10; num++ {
			if newNum := complexDeduceCoordinates(set, board, y, x); newNum != 0{
				return newNum
			}
		}
	}
	return 0
}

func complexDeduceCoordinates(coordinates [][]int, board [][]int, y int, x int)int{
	possibles := [][]int{}
	for _, v := range coordinates {
		if board[v[0]][v[1]] == 0 && (v[0] != y || v[1] != x){
			possibles = append(possibles, findPossibles(board, v[0], v[1]))
		}
	}
	for _, testNum := range findPossibles(board, y, x){
		isElsewhere := false
		for _, pNums := range possibles {
			if intContains(pNums, testNum) {
				isElsewhere = true
				break
			}
		}
		if !isElsewhere{
			return testNum
		}
	}
	return 0
}


func properSolveBoard(board [][]int)bool{
	for !isSolved(board){
		simpleSolutions := 0
		for y:=0;y<9;y++{
			for x:=0;x<9;x++{
				if board[y][x] == 0 {
					newValue := simpleDeduce(board, y, x)
					if newValue != 0 {
						board[y][x] = newValue
						simpleSolutions++
					}
				}
			}
		}
		if simpleSolutions == 0{
			for y:=0;y<9;y++ {
				for x := 0; x < 9; x++ {
					if board[y][x] == 0 {
						newValue := complexDeduce(board, y, x)
						if newValue != 0 {
							board[y][x] = newValue
							simpleSolutions++
						}
					}
				}
			}
			return false
		}
	}
	return true
}

func isSolved(board [][]int)bool{
	for r := 0; r < 9; r++ {
		row, col, box := 0, 0, 0
		for c := 0; c < 9; c++ {
			i := (r % 3) * 3 + c % 3
			j := (r / 3) * 3 + c / 3
			row ^= 1 << uint(board[r][c])
			col ^= 1 << uint(board[c][r])
			box ^= 1 << uint(board[j][i])
		}
		if row != 1022 || col != 1022 || box != 1022 {
			return false
		}
	}
	return true
}

func printBoard(board [][]int){
	out1  := ""
	for i1, v := range board {
		out1 += strings.ReplaceAll(" " + strconv.Itoa(v[0]) + "  " + strconv.Itoa(v[1]) + "  " + strconv.Itoa(v[2]) + " | " + strconv.Itoa(v[3]) + "  " + strconv.Itoa(v[4]) + "  " + strconv.Itoa(v[5]) + " | " + strconv.Itoa(v[6]) + "  " + strconv.Itoa(v[7]) + "  " + strconv.Itoa(v[8]), "0", " ") + "\n"
		if i1%3 == 2 && i1 != 8{
			out1 +=" ----------------------------\n"
		}
	}
	fmt.Println(out1)
}

func findPossibles(board [][]int, y int, x int)[]int{
	pool := []int{}
	for i:=1;i<10;i++{
		virt := []int{board[0][x], board[1][x], board[2][x], board[3][x], board[4][x], board[5][x], board[6][x], board[7][x], board[8][x]}
		box := []int{board[(y/3)*3 + 0][(x/3)*3 + 0], board[(y/3)*3 + 1][(x/3)*3 + 0], board[(y/3)*3 + 2][(x/3)*3 + 0], board[(y/3)*3 + 0][(x/3)*3 + 1], board[(y/3)*3 + 1][(x/3)*3 + 1], board[(y/3)*3 + 2][(x/3)*3 + 1], board[(y/3)*3 + 0][(x/3)*3 + 2], board[(y/3)*3 + 1][(x/3)*3 + 2], board[(y/3)*3 + 2][(x/3)*3 + 2]}
		total := countNumbers(board[y], i) + countNumbers(virt, i) + countNumbers(box, i)
		if total == 0{
			pool = append(pool, i)
		}
	}
	return pool
}

func countNumbers(arr []int, findNum int)int{
	total := 0
	for _, v := range arr{
		if v == findNum{
			total++
		}
	}
	return total
}

func main() {

	var board = [][]int{
		[]int{8,5,0,0,0,0,0,0,0},
		[]int{0,0,0,1,3,0,0,0,0},
		[]int{4,0,0,0,0,0,0,9,5},
		[]int{0,2,0,9,0,8,0,0,0},
		[]int{5,8,1,0,0,0,0,0,0},
		[]int{0,0,0,0,0,0,7,0,6},
		[]int{0,0,7,0,0,1,8,0,0},
		[]int{0,0,0,8,0,9,0,0,0},
		[]int{0,0,0,0,0,0,2,1,3},
	}

	start := time.Now()
	solved := properSolveBoard(board)
	if !solved {
		solved = longShot(board)
	}

	defer fmt.Println(time.Since(start))
	if solved{
		defer fmt.Println("\nComplete!")
	}else{
		defer fmt.Println("\nI got stuck...")
	}
	printBoard(board)

}
