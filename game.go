package main

import (
	"image/color"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 350
	screenHeight = 350
	boardSize    = 3
	cellSize     = screenWidth / boardSize
	lineWidth    = 2
)

type Player int

const (
	Empty Player = iota
	Circle
	Cross
)

var (
	board        [boardSize][boardSize]Player
	currentTurn  = Circle
	winner       = Empty
	winnerString string
	gameOver     = false
	highlighted  = false
	bestMoveRow  = -1
	bestMoveCol  = -1
)

type MoveScore struct {
	move  [2]int
	score int
}

func FindBestMove() {
	bestVal := math.MinInt32
	var moveSequence []MoveScore

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == Empty {
				board[i][j] = Circle
				moveVal := minimax(board, 0, false, &moveSequence)
				board[i][j] = Empty

				if moveVal > bestVal {
					bestMoveRow, bestMoveCol = i, j
					bestVal = moveVal
				}
			}
		}
	}

}

func minimax(board [boardSize][boardSize]Player, depth int, isMaximizing bool, moveSequence *[]MoveScore) int {
	score := evaluate(board)

	if score != -2 {
		return score
	}

	if isMaximizing {
		best := math.MinInt32

		for i := 0; i < boardSize; i++ {
			for j := 0; j < boardSize; j++ {
				if board[i][j] == Empty {
					board[i][j] = Circle
					moveVal := minimax(board, depth+1, !isMaximizing, moveSequence)
					board[i][j] = Empty

					moveScore := MoveScore{
						move:  [2]int{j, i},
						score: moveVal,
					}
					*moveSequence = append(*moveSequence, moveScore)

					best = max(best, moveVal)
				}
			}
		}

		return best
	} else {
		best := math.MaxInt32

		for i := 0; i < boardSize; i++ {
			for j := 0; j < boardSize; j++ {
				if board[i][j] == Empty {
					board[i][j] = Cross
					moveVal := minimax(board, depth+1, !isMaximizing, moveSequence)
					board[i][j] = Empty

					moveScore := MoveScore{
						move:  [2]int{j, i},
						score: moveVal,
					}
					*moveSequence = append(*moveSequence, moveScore)

					best = min(best, moveVal)
				}
			}
		}

		return best
	}

}

func evaluate(board [boardSize][boardSize]Player) int {
	if checkWin(board, Cross) {
		return -1
	} else if checkWin(board, Circle) {
		return 1
	}

	isBoardFull := true
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == Empty {
				isBoardFull = false
				break
			}
		}
	}

	if isBoardFull {
		return 0 // Ничья
	}

	return -2
}

func checkWin(board [boardSize][boardSize]Player, player Player) bool {
	for i := 0; i < boardSize; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}

	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func isBoardFull() bool {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == Empty {
				return false
			}
		}
	}
	return true
}

func checkWinner() Player {
	for i := 0; i < boardSize; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != Empty {
			return board[i][0]
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != Empty {
			return board[0][i]
		}
	}

	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != Empty {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != Empty {
		return board[0][2]
	}

	return Empty
}

func resetGame() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board[i][j] = Empty
		}
	}

	currentTurn = Cross
	winner = Empty
	winnerString = ""
	gameOver = false
	highlighted = false
	bestMoveRow = -1
	bestMoveCol = -1
}

func update(screen *ebiten.Image) error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !gameOver {
		x, y := ebiten.CursorPosition()
		if x < 0 || y < 0 || x >= screenWidth || y >= screenHeight {
			return nil
		}

		row, col := y/cellSize, x/cellSize
		if board[row][col] == Empty {
			board[row][col] = currentTurn
			winner = checkWinner()
			if winner != Empty || isBoardFull() {
				if winner != Empty {
					switch winner {
					case Circle:
						winnerString = "Player 1 wins"
					case Cross:
						winnerString = "Player 2 wins"
					}
				} else {
					winnerString = "It's a draw!"
				}
				gameOver = true
			} else {
				if currentTurn == Circle {
					currentTurn = Cross
				} else {
					currentTurn = Circle
				}
			}
		}
	}

	if !gameOver && currentTurn == Circle {
		FindBestMove()
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		resetGame()
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	for i := 1; i < boardSize; i++ {
		ebitenutil.DrawLine(screen, 0, float64(i*cellSize), screenWidth, float64(i*cellSize), color.Black)
		ebitenutil.DrawLine(screen, float64(i*cellSize), 0, float64(i*cellSize), screenHeight, color.Black)
	}

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			var symbol string
			var textColor color.Color
			switch board[i][j] {
			case Circle:
				symbol = "O"
				textColor = color.RGBA{36, 36, 36, 255}
			case Cross:
				symbol = "X"
				textColor = color.RGBA{65, 65, 65, 255}
			default:
				textColor = color.White
			}

			if i == bestMoveRow && j == bestMoveCol && currentTurn == Circle {
				textColor = color.RGBA{255, 0, 0, 255}
			}

			ebitenutil.DrawRect(screen, float64(j*cellSize)+lineWidth, float64(i*cellSize)+lineWidth, float64(cellSize)-2*lineWidth, float64(cellSize)-2*lineWidth, textColor)

			if symbol != "" {
				ebitenutil.DebugPrintAt(screen, symbol, j*cellSize+cellSize/2-5, i*cellSize+cellSize/2-10)
			}
		}
	}

	if gameOver {
		bgColor := color.RGBA{255, 0, 0, 255}
		ebitenutil.DrawRect(screen, 0, 0, screenWidth, 20, bgColor)
		ebitenutil.DebugPrintAt(screen, winnerString+"  (R-reset; Q-exit)", 75, 0)
	}

	return nil
}
