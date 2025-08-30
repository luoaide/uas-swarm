package swarm

// This is just for thinking.
// Leadership is a binary tree, heavy high and left w/ C nodes (inherantly higher priority)

// I don't care about "leadership" I care about who executes which tasks

type Node struct {
	Value  int
	parent *Node
	left   *Node
	right  *Node
}

func (n *Node) Promise() {
	// MY LEFT IS n.left AND MY RIGHT IS n.right
	// go broadcast that promise out the the c network and the radio network... leave this goroutine hanging
	// when left receives, send left's promise.. I receive left's promise
}
