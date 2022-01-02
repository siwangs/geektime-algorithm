type MyCircularDeque struct {
	Size int
	Cap  int
	Head *Node
	Tail *Node
}

type Node struct {
	Val  int
	Pre  *Node
	Next *Node
}

func Constructor(k int) MyCircularDeque {
	//Initial head and tail node
	head := &Node{-1, nil, nil}
	tail := &Node{-1, nil, nil}

	//Initial linked list and return
	head.Next = tail
	tail.Pre = head

	return MyCircularDeque{0, k, head, tail}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.Size < this.Cap {
		// set up a list node, insert if Size < Cap
		cur := &Node{Val: value, Pre: this.Head, Next: this.Head.Next}
		this.Head.Next.Pre = cur
		this.Head.Next = cur

		this.Size++
		return true
	}
	return false
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.Size < this.Cap {
		curr := &Node{value, this.Tail.Pre, this.Tail}
		this.Tail.Pre.Next = curr
		this.Tail.Pre = curr
		this.Size++
		return true
	}
	return false
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.Size > 0 {
		this.Head.Next.Next.Pre = this.Head
		this.Head.Next = this.Head.Next.Next
		this.Size--
		return true
	}
	return false
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.Size > 0 {
		this.Tail.Pre.Pre.Next = this.Tail
		this.Tail.Pre = this.Tail.Pre.Pre
		this.Size--
		return true
	}
	return false
}

func (this *MyCircularDeque) GetFront() int {
	if this.Size > 0 {
		return this.Head.Next.Val
	}
	return -1
}

func (this *MyCircularDeque) GetRear() int {
	if this.Size > 0 {
		return this.Tail.Pre.Val
	}
	return -1
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.Size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.Size == this.Cap
}
