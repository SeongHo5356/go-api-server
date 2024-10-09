package lab

import "fmt"

type Node struct{
	Left *Node
	Right *Node
	Value int	
}

// 새로운 노드를 생성
func MakeNode(value int) *Node{
	return &Node{
		Left: nil,
		Right: nil,
		Value: value,
	}
}

// 트리에 노드를 추가 및 정렬
func InsertNode(root *Node, value int) *Node{
	if root == nil{
		return MakeNode(value)
	}

	if value < root.Value{
		root.Left = InsertNode(root.Left, value)
	}else if value > root.Value{ 
		root.Right = InsertNode(root.Right, value)
	}

	return root
}

// 중위 순회 (왼쪽->루트->오른쪽)
func InOrder(root *Node){
	if root != nil{
		InOrder(root.Left)
		fmt.Printf("%d ", root.Value)
		InOrder(root.Right)
	}
}

func BinarySearchTree(){
	var root *Node

	// 노드 추가
    root = InsertNode(root, 10)
    root = InsertNode(root, 5)
    root = InsertNode(root, 20)
    root = InsertNode(root, 1)
    root = InsertNode(root, 8)
    
    // 중위 순회
    fmt.Print("InOrder traversal: ")
    InOrder(root)
    fmt.Println()
}
