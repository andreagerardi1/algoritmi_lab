package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList *Node

func newNode(value int) *Node {
	return &Node{value, nil}
}

// aggiunge un nodo in testa
func addNewNode(l *LinkedList, value int) {
	node := newNode(value)
	node.next = *l
	*l = node
}

func searchList(l LinkedList, val int) (isIn bool, nodo *Node) {
	p := l
	for p != nil {
		if p.value == val {
			isIn = true
			nodo = p
			break
		}
		p = p.next
	}
	return
}

func deleteItem(l LinkedList, val int) bool {
	var prev *Node
	cur := l
	for cur != nil {
		if cur.value == val {
			if prev == nil {
				l = l.next
			} else {
				prev.next = cur.next
			}
			return true
		}
		prev = cur
		cur = cur.next
	}
	return false
}

// stampa la lista non vuota
func printList(l LinkedList) {
	p := l //p serve per non cambiare la lista
	for p != nil {
		fmt.Print(p.value, " ")
		p = p.next
	}
	fmt.Println()
}

func main() {
	var l LinkedList
	fmt.Println("l", l)
	addNewNode(&l, 10)
	//fmt.Println("l dopo il primo", l.value)
	addNewNode(&l, 20)
	//fmt.Println("l dopo il primo", l.value)
	addNewNode(&l, 30)
	//fmt.Println("l dopo il primo", l.value)
	f, n := searchList(l, 10)
	if f == false {
		fmt.Println("element is not present")
	} else {
		fmt.Println("10 is present in node: ", n)
	}
	b := deleteItem(l, 10)
	fmt.Println(b)
	printList(l)
}
