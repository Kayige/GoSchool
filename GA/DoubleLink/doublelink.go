package main

import (
	"errors"
	"fmt"
)

//Node definition in struct
type Node struct {
	item string
	prev *Node
	next *Node
}

type dblinklist struct {
	tail *Node
	head *Node
	size int
}

func (p *dblinklist) addNode(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
		prev: nil,
	}
	if p.head == nil {
		p.head = newNode
		p.tail = newNode
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		p.tail = newNode
	}
	p.size++
	return nil
}

func (p *dblinklist) printAllNodes() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("DB Link list is empty")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}

	return nil
}

func (p *dblinklist) printAllNodesReverse() error {
	currentNode := p.tail
	if currentNode == nil {
		fmt.Println("DB Link list is empty")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.prev != nil {
		fmt.Printf("%+v\n", currentNode.prev.item)
		currentNode = currentNode.prev
	}
	return nil
}

func (p *dblinklist) addAtPos(index int, name string) error {
	newNode := &Node{
		item: name,
		next: nil,
		prev: nil,
	}
	if index == 1 {
		p.head.prev = newNode
		newNode.next = p.head
		p.head = newNode
		return nil

	} else if index > 0 && index <= p.size+1 {
		currentNode := p.head
		var prevNode *Node
		for i := 1; i <= index-1; i++ {
			prevNode = currentNode
			currentNode = currentNode.next
		}
		currentNode.prev = newNode
		newNode.next = currentNode
		prevNode.next = newNode
		newNode.prev = prevNode
		return nil
	}
	return errors.New("Invalid Index")
}

func (p *dblinklist) removeHead() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("DB Link list is empty")
		return nil
	}

	if currentNode.next != nil {
		p.head = currentNode.next
		p.head.prev = nil
	}
	p.size--
	return nil

}

func (p *dblinklist) removeTail() error {
	currentNode := p.tail

	if currentNode == nil {
		fmt.Println("DB Link list is empty")
		return nil
	}
	if currentNode.prev != nil {
		p.tail = currentNode.prev
		p.tail.next = nil
	}
	p.size--
	return nil
}

func (p *dblinklist) remove(index int) error {
	if p.head == nil {
		return errors.New("empty double linked list")
	}
	if index > 0 && index <= p.size {
		if index == 1 {
			p.removeHead()
		} else {
			var currentNode *Node = p.head
			var prevNode *Node
			for i := 1; i <= index-1; i++ {
				prevNode = currentNode
				currentNode = currentNode.next

			}
			prevNode.next = currentNode.next
			currentNode.next.prev = prevNode
		}
	}
	p.size--
	return nil
}

func main() {
	dblist := &dblinklist{}
	dblist.addNode("Mary")
	dblist.addNode("Joe")
	dblist.addNode("David")
	dblist.printAllNodes()
	fmt.Println()
	dblist.printAllNodesReverse()
	dblist.addAtPos(2, "NewNode")
	fmt.Println()
	dblist.printAllNodes()
	dblist.remove(3)
	fmt.Println()
	dblist.printAllNodes()

}
