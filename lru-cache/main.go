package main

import "fmt"

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Hash map[string]*Node

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{
		Queue: newQueue(), Hash: Hash{},
	}
}

func newQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{
		Head: head, Tail: tail,
	}
}

func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.Val)
	left := n.Left
	right := n.Right

	right.Left = left
	left.Right = right
	c.Queue.Length--
	delete(c.Hash, n.Val)
	return n
}

func (c *Cache) Add(n *Node) {
	t := c.Queue.Head.Right
	c.Queue.Head.Right = n
	n.Right = t
	t.Left = n

	c.Queue.Length++
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.Queue.display()
}

func (q *Queue) display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf(" <==> ")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("Starting cache...")

	cache := NewCache()
	for _, word := range []string{"parrot", "avocado", "tree", "cat"} {
		cache.Check(word)
		cache.Display()
	}
}
