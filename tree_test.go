package main

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestAVLTree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Тест для создания пустого дерева и вставки элементов.
	t.Run("Insert", func(t *testing.T) {
		tree := NewAVLTree(10)
		values := []int{5, 15, 3, 7, 12, 17}

		for _, v := range values {
			tree.Insert(v)
		}

		inOrder := InOrderTraversal(tree.Root)
		expected := []int{3, 5, 7, 10, 12, 15, 17}

		if !reflect.DeepEqual(inOrder, expected) {
			t.Errorf("Ожидался порядок %v, но получен %v", expected, inOrder)
		}
	})

	// Тест для проверки балансировки дерева.
	t.Run("Balance", func(t *testing.T) {
		tree := NewAVLTree(10)
		values := []int{5, 15, 3, 7, 12, 17}

		for _, v := range values {
			tree.Insert(v)
		}

		// Проверяем балансировку дерева после вставки элементов.
		balanceFactors := []int{
			balanceFactor(tree.Root),
			balanceFactor(tree.Root.Left),
			balanceFactor(tree.Root.Right),
			balanceFactor(tree.Root.Left.Left),
			balanceFactor(tree.Root.Left.Right),
			balanceFactor(tree.Root.Right.Left),
			balanceFactor(tree.Root.Right.Right),
		}

		expected := []int{0, 0, 0, 0, 0, 0, 0}
		if !reflect.DeepEqual(balanceFactors, expected) {
			t.Errorf("Ожидался баланс %v, но получен %v", expected, balanceFactors)
		}
	})

	// Тест для проверки вставки дублирующихся элементов.
	t.Run("DuplicateInsert", func(t *testing.T) {
		tree := NewAVLTree(10)
		values := []int{5, 15, 3, 7, 12, 17, 10}

		for _, v := range values {
			tree.Insert(v)
		}

		// Ожидается, что дублирующиеся значения будут проигнорированы.
		inOrder := InOrderTraversal(tree.Root)
		expected := []int{3, 5, 7, 10, 12, 15, 17}

		if !reflect.DeepEqual(inOrder, expected) {
			t.Errorf("Ожидался порядок %v, но получен %v", expected, inOrder)
		}
	})

	// Тест для удаления элемента из дерева.
	t.Run("Delete", func(t *testing.T) {
		tree := NewAVLTree(10)
		values := []int{5, 15, 3, 7, 12, 17}

		for _, v := range values {
			tree.Insert(v)
		}

		// Удаляем элемент 12 из дерева.
		tree.Delete(12)

		inOrder := InOrderTraversal(tree.Root)
		expected := []int{3, 5, 7, 10, 15, 17}

		if !reflect.DeepEqual(inOrder, expected) {
			t.Errorf("Ожидался порядок %v, но получен %v", expected, inOrder)
		}
	})

	// Тест для удаления корневого элемента из дерева.
	t.Run("DeleteRoot", func(t *testing.T) {
		tree := NewAVLTree(10)
		values := []int{5, 15, 3, 7, 12, 17}

		for _, v := range values {
			tree.Insert(v)
		}

		// Удаляем корневой элемент 10 из дерева.
		tree.Delete(10)

		inOrder := InOrderTraversal(tree.Root)
		expected := []int{3, 5, 7, 12, 15, 17}

		if !reflect.DeepEqual(inOrder, expected) {
			t.Errorf("Ожидался порядок %v, но получен %v", expected, inOrder)
		}
	})

	// Тест для удаления элемента, отсутствующего в дереве.
	t.Run("DeleteNonExistent", func(t *testing.T) {
		tree := NewAVLTree(10)
		values := []int{5, 15, 3, 7, 12, 17}

		for _, v := range values {
			tree.Insert(v)
		}

		// Пытаемся удалить элемент 8, который отсутствует в дереве.
		tree.Delete(8)

		inOrder := InOrderTraversal(tree.Root)
		expected := []int{3, 5, 7, 10, 12, 15, 17}

		if !reflect.DeepEqual(inOrder, expected) {
			t.Errorf("Ожидался порядок %v, но получен %v", expected, inOrder)
		}
	})

	// Тест для удаления всех элементов из дерева.
	t.Run("DeleteAll", func(t *testing.T) {
		tree := NewAVLTree(10)
		values := []int{5, 15, 3, 7, 12, 17}

		for _, v := range values {
			tree.Insert(v)
		}

		// Удаляем все элементы из дерева.
		for _, v := range values {
			tree.Delete(v)
		}

		inOrder := InOrderTraversal(tree.Root)
		expected := []int{10}

		if !reflect.DeepEqual(inOrder, expected) {
			t.Errorf("Ожидался порядок %v, но получен %v", expected, inOrder)
		}
	})

	// Тест для удаления корня с двумя детьми.
	t.Run("DeleteRootWithTwoChildren", func(t *testing.T) {
		tree := NewAVLTree(10)
		values := []int{5, 15, 3, 7, 12, 17, 8}

		for _, v := range values {
			tree.Insert(v)
		}

		// Удаляем корневой элемент 10 из дерева.
		tree.Delete(10)

		inOrder := InOrderTraversal(tree.Root)
		expected := []int{3, 5, 7, 8, 12, 15, 17}

		if !reflect.DeepEqual(inOrder, expected) {
			t.Errorf("Ожидался порядок %v, но получен %v", expected, inOrder)
		}
	})
}

func TestRandomInsertDelete(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Тест для случайной вставки и удаления элементов.
	t.Run("RandomInsertDelete", func(t *testing.T) {
		tree := NewAVLTree(50)
		var inserted []int

		for i := 0; i < 1000; i++ {
			value := rand.Intn(100)
			op := rand.Intn(2)

			if op == 0 {
				tree.Insert(value)
				inserted = append(inserted, value)
			} else if len(inserted) > 0 {
				index := rand.Intn(len(inserted))
				valueToDelete := inserted[index]
				tree.Delete(valueToDelete)
				inserted = append(inserted[:index], inserted[index+1:]...)
			}

			// Проверяем, что дерево всегда остается сбалансированным.
			checkBalance(t, tree.Root)
		}

		// Проверяем, что элементы в дереве соответствуют вставленным.
		inOrder := InOrderTraversal(tree.Root)
		expected := make([]int, len(inserted))
		copy(expected, inserted)
		sortIntSlice(expected)

		if !reflect.DeepEqual(inOrder, expected) {
			t.Errorf("Ожидался порядок %v, но получен %v", expected, inOrder)
		}
	})
}

func sortIntSlice(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

func checkBalance(t *testing.T, root *Node) int {
	if root == nil {
		return 0
	}

	leftHeight := checkBalance(t, root.Left)
	rightHeight := checkBalance(t, root.Right)

	balance := leftHeight - rightHeight
	if balance > 1 || balance < -1 {
		t.Errorf("Несбалансированное поддерево с корнем %d, баланс: %d", root.Value, balance)
	}

	return maxB(leftHeight, rightHeight) + 1
}

func TestMain(m *testing.M) {
	fmt.Println("Запуск тестов...")
	code := m.Run()
	fmt.Println("Завершение тестов.")
	os.Exit(code)
}
