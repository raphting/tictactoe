package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func createEmptyBoard() [][]int {
	// Create empty abstract 3x3 board
	// 0 = not set; 1 = player one; 2 = player two
	fields := make([][]int, 3)
	for i := range fields {
		fields[i] = make([]int, 3)
		for empty := range fields[i] {
			fields[i][empty] = 0
		}
	}
	return fields
}

func getBoard(board [][]int) string {
	boardDesign := []string{"-------", "|1|2|3|", "-------", "|4|5|6|", "-------", "|7|8|9|", "-------"}

	// Render every choice from the abstract board into the board design
	for x := range board {
		for y := range board[x] {
			// No choice yet for this field
			if board[x][y] == 0 {
				continue
			}

			// Choose which stone to play (player one = X; player two = O)
			place := "X"
			if board[x][y] == 2 {
				place = "O"
			}

			// Calculate field number and place stone
			field := strconv.Itoa(1 + (y + (3 * x)))
			for key, line := range boardDesign {
				boardDesign[key] = strings.Replace(line, field, place, 1)
			}
		}
	}

	return strings.Join(boardDesign, "\n")
}

// Parse input and return x, y coordinates for abstract board
func calculateField(input string) (int, int, error) {
	num, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, 0, err
	}

	if num < 1 || num > 9 {
		return 0, 0, errors.New("Field out of board dimension")
	}

	if num <= 3 {
		return 0, num - 1, nil
	}

	if num <= 6 {
		return 1, num - 4, nil
	}

	return 2, num - 7, nil
}

// Ask user for input
func actionRequired(player int) string {
	name := "one"
	if player == 1 {
		name = "two"
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Player " + name + ", enter field number: ")
	text, _ := reader.ReadString('\n')
	return text
}

func whoWon(board [][]int) int {
	// Check horizontal
	for line := range board {
		if board[line][0] == 0 {
			continue
		}
		if board[line][0] == board[line][1] && board[line][1] == board[line][2] {
			return board[line][0]
		}
	}

	// Check vertical
	for i := 0; i < 3; i++ {
		if board[0][i] == 0 {
			continue
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}

	// Check diagonal
	// If stone in middle is not set, there is no diagonal winner
	if board[1][1] == 0 {
		return 0
	}

	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[1][1]
	}

	if board[2][0] == board[1][1] && board[1][1] == board[0][2] {
		return board[1][1]
	}

	return 0
}

func computerInput(board [][]int) (int, int) {
	empty := make([][]int, 0)

	for x := range board {
		for y := range board[x] {
			if board[x][y] == 0 {
				i := []int{x, y}
				empty = append(empty, i)
			}
		}
	}

	field := rand.Intn(len(empty))
	return empty[field][0], empty[field][1]
}

func startEventLoop(board [][]int) {
	player := 0
	for {
		x, y := 0, 0

		if player%2 == 1 {
			x, y = computerInput(board)
		} else {
			field := actionRequired(0)
			var err error
			x, y, err = calculateField(field)
			if err != nil {
				fmt.Println("Whoops, your input was not a natural number between 1 and 9. Try again!")
				continue
			}
		}

		if board[x][y] != 0 {
			fmt.Println("Please choose an empty field to place your stone. Try again!")
			continue
		}

		board[x][y] = player%2 + 1

		fmt.Println(getBoard(board))

		winner := whoWon(board)
		if winner > 0 {
			name := "Player one"
			if winner == 2 {
				name = "Computer"
			}
			fmt.Println(name + " won.")
			break
		}

		// After 9 rounds, the game is over, even without a winner
		player++
		if player > 8 {
			break
		}
	}
}

func main() {
	fmt.Println("Welcome to the Game of Tic Tac Toe")
	board := createEmptyBoard()
	fmt.Println(getBoard(board))

	startEventLoop(board)
}
