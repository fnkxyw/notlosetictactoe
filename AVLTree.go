package main

// Node представляет узел бинарного дерева.
type Node struct {
	Value       int
	Left, Right *Node
	Height      int
}

// AVLTree представляет бинарное сбалансированное дерево.
type AVLTree struct {
	Root *Node
}

func deleteNode(root *Node, value int) *Node {
	if root == nil {
		return root
	}

	// Рекурсивно ищем узел для удаления в левом или правом поддереве.
	if value < root.Value {
		root.Left = deleteNode(root.Left, value)
	} else if value > root.Value {
		root.Right = deleteNode(root.Right, value)
	} else {
		// Узел с заданным значением найден и будет удален.

		// Узел без левого или правого поддерева.
		if root.Left == nil || root.Right == nil {
			var temp *Node
			if temp = root.Left; temp == nil {
				temp = root.Right
			}

			// Если у узла нет детей, просто удаляем его.
			if temp == nil {
				temp = root
				root = nil
			} else { // Копируем данные из непустого узла.
				*root = *temp
			}
			temp = nil
		} else { // Узел имеет оба поддерева, найдем наименьший узел в правом поддереве (или наибольший в левом).
			temp := minValueNode(root.Right)
			root.Value = temp.Value
			root.Right = deleteNode(root.Right, temp.Value)
		}
	}

	if root == nil {
		return root
	}

	// Обновляем высоту текущего узла.
	updateHeight(root)

	// Получаем баланс текущего узла и выполняем балансировку.
	balance := balanceFactor(root)

	// Левое поддерево перекосило влево.
	if balance > 1 {
		if balanceFactor(root.Left) >= 0 {
			return rotateRight(root)
		} else {
			root.Left = rotateLeft(root.Left)
			return rotateRight(root)
		}
	}

	// Правое поддерево перекосило вправо.
	if balance < -1 {
		if balanceFactor(root.Right) <= 0 {
			return rotateLeft(root)
		} else {
			root.Right = rotateRight(root.Right)
			return rotateLeft(root)
		}
	}

	return root
}

// Delete удаляет узел с заданным значением из дерева.
func (t *AVLTree) Delete(value int) {
	t.Root = deleteNode(t.Root, value)
}

// minValueNode находит узел с минимальным значением в дереве.
func minValueNode(root *Node) *Node {
	current := root
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// NewNode создает новый узел с заданным значением.
func NewNode(value int) *Node {
	return &Node{Value: value, Height: 1}
}

// NewAVLTree создает новое бинарное сбалансированное дерево с заданным корнем.
func NewAVLTree(rootValue int) *AVLTree {
	return &AVLTree{Root: NewNode(rootValue)}
}

// height возвращает высоту узла.
func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.Height
}

// max возвращает максимальное из двух чисел.
func maxB(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// balanceFactor возвращает разницу высот поддеревьев.
func balanceFactor(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.Left) - height(n.Right)
}

// updateHeight обновляет высоту узла.
func updateHeight(n *Node) {
	if n == nil {
		return
	}
	n.Height = 1 + maxB(height(n.Left), height(n.Right))
}

// rotateRight выполняет правый поворот узла.
func rotateRight(y *Node) *Node {
	x := y.Left
	y.Left = x.Right
	x.Right = y
	updateHeight(y)
	updateHeight(x)
	return x
}

// rotateLeft выполняет левый поворот узла.
func rotateLeft(x *Node) *Node {
	y := x.Right
	x.Right = y.Left
	y.Left = x
	updateHeight(x)
	updateHeight(y)
	return y
}

// insert вставляет новый элемент в дерево.
func insert(root *Node, value int) *Node {
	if root == nil {
		return NewNode(value)
	}

	if value < root.Value {
		root.Left = insert(root.Left, value)
	} else if value > root.Value {
		root.Right = insert(root.Right, value)
	} else {
		// Дублирующиеся значения не допускаются (по вашему выбору).
		return root
	}

	updateHeight(root)
	balance := balanceFactor(root)

	// Проверяем и выполняем балансировку узла, если необходимо.
	if balance > 1 {
		if value < root.Left.Value {
			return rotateRight(root)
		} else {
			root.Left = rotateLeft(root.Left)
			return rotateRight(root)
		}
	}
	if balance < -1 {
		if value > root.Right.Value {
			return rotateLeft(root)
		} else {
			root.Right = rotateRight(root.Right)
			return rotateLeft(root)
		}
	}

	return root
}

// Insert добавляет новый элемент в дерево.
func (t *AVLTree) Insert(value int) {
	t.Root = insert(t.Root, value)
}

// InOrderTraversal выполняет обход в порядке возрастания и возвращает значения узлов в виде среза.
func InOrderTraversal(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}
	result = append(result, InOrderTraversal(root.Left)...)
	result = append(result, root.Value)
	result = append(result, InOrderTraversal(root.Right)...)
	return result
}
