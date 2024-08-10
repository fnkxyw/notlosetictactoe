package main

import (
	"testing"
)

func TestFindBestMove(t *testing.T) {
	// Создаем тестовое игровое поле

	// Заполняем поле, чтобы убедиться, что bestMoveRow и bestMoveCol изначально равны -1
	bestMoveRow = -1
	bestMoveCol = -1

	// Вызываем функцию FindBestMove
	FindBestMove()

	// Проверяем, что bestMoveRow и bestMoveCol теперь имеют правильные значения
	expectedRow := 0
	expectedCol := 0
	if bestMoveRow != expectedRow || bestMoveCol != expectedCol {
		t.Errorf("Ожидалось (%d, %d), но получено (%d, %d)", expectedRow, expectedCol, bestMoveRow, bestMoveCol)
	}

	// Сбрасываем значения bestMoveRow и bestMoveCol
	bestMoveRow = -1
	bestMoveCol = -1

	// Создаем другое тестовое игровое поле
	board = [3][3]Player{
		{Circle, Empty, Cross},
		{Cross, Circle, Cross},
		{Empty, Circle, Empty},
	}

	// Вызываем функцию FindBestMove
	FindBestMove()

	// Проверяем, что bestMoveRow и bestMoveCol теперь имеют правильные значения
	expectedRow = 0
	expectedCol = 1
	if bestMoveRow != expectedRow || bestMoveCol != expectedCol {
		t.Errorf("Ожидалось (%d, %d), но получено (%d, %d)", expectedRow, expectedCol, bestMoveRow, bestMoveCol)
	}
}

func TestEvaluate(t *testing.T) {
	// Создаем тестовое игровое поле
	board := [3][3]Player{
		{Circle, Circle, Cross},
		{Cross, Cross, Circle},
		{Circle, Cross, Circle},
	}

	// Вызываем функцию evaluate и проверяем, что она возвращает ожидаемое значение.
	expected := 0
	result := evaluate(board)
	if result != expected {
		t.Errorf("Ожидалось %d, но получено %d", expected, result)
	}

	// Создаем еще одно тестовое поле
	board = [3][3]Player{
		{Circle, Cross, Cross},
		{Cross, Circle, Circle},
		{Circle, Cross, Cross},
	}

	// Вызываем функцию evaluate и проверяем, что она возвращает ожидаемое значение.
	expected = 0
	result = evaluate(board)
	if result != expected {
		t.Errorf("Ожидалось %d, но получено %d", expected, result)
	}

	// Создаем поле для ничьей
	board = [3][3]Player{
		{Circle, Cross, Circle},
		{Circle, Circle, Cross},
		{Cross, Circle, Cross},
	}

	// Вызываем функцию evaluate и проверяем, что она возвращает ожидаемое значение.
	expected = 0
	result = evaluate(board)
	if result != expected {
		t.Errorf("Ожидалось %d, но получено %d", expected, result)
	}
}

func TestCheckWin(t *testing.T) {
	board := [3][3]Player{
		{Empty, Cross, Circle},
		{Cross, Cross, Circle},
		{Cross, Circle, Empty},
	}

	// Вызываем функцию checkWin для крестиков и проверяем, что она возвращает true.
	expected := false
	result := checkWin(board, Cross)
	if result != expected {
		t.Errorf("Ожидалось %v, но получено %v", expected, result)
	}

	// Создаем тестовое игровое поле для выигрыша ноликов
	board = [3][3]Player{
		{Circle, Circle, Circle},
		{Cross, Cross, Empty},
		{Empty, Empty, Empty},
	}

	// Вызываем функцию checkWin для ноликов и проверяем, что она возвращает true.
	expected = true
	result = checkWin(board, Circle)
	if result != expected {
		t.Errorf("Ожидалось %v, но получено %v", expected, result)
	}

	// Создаем тестовое игровое поле без выигрышных комбинаций
	board = [3][3]Player{
		{Circle, Cross, Circle},
		{Cross, Circle, Cross},
		{Circle, Cross, Circle},
	}

	// Вызываем функцию checkWin для крестиков и ноликов и проверяем, что она возвращает false.
	expected = false
	result = checkWin(board, Cross)
	if result != expected {
		t.Errorf("Ожидалось %v, но получено %v", expected, result)
	}
	expected = true
	result = checkWin(board, Circle)
	if result != expected {
		t.Errorf("Ожидалось %v, но получено %v", expected, result)
	}
}

func TestMax(t *testing.T) {
	// Вызываем функцию max с разными значениями x и y и проверяем, что она возвращает максимальное значение.
	x := 5
	y := 10
	expected := 10
	result := max(x, y)
	if result != expected {
		t.Errorf("Ожидалось %d, но получено %d", expected, result)
	}

	x = 15
	y = 8
	expected = 15
	result = max(x, y)
	if result != expected {
		t.Errorf("Ожидалось %d, но получено %d", expected, result)
	}
}

func TestResetGame(t *testing.T) {
	// Устанавливаем некоторые значения в переменных
	currentTurn = Cross
	winner = Circle
	gameOver = true
	bestMoveRow = 1
	bestMoveCol = 2

	// Вызываем функцию resetGame
	resetGame()

	// Проверяем, что все переменные сброшены к начальным значениям
	if currentTurn != Cross || winner != Empty || gameOver || bestMoveRow != -1 || bestMoveCol != -1 {
		t.Errorf("Переменные не были сброшены к начальным значениям")
	}
}

func TestMin(t *testing.T) {
	// Вызываем функцию min с разными значениями x и y и проверяем, что она возвращает минимальное значение.
	x := 5
	y := 10
	expected := 5
	result := min(x, y)
	if result != expected {
		t.Errorf("Ожидалось %d, но получено %d", expected, result)
	}

	x = 15
	y = 8
	expected = 8
	result = min(x, y)
	if result != expected {
		t.Errorf("Ожидалось %d, но получено %d", expected, result)
	}
}
