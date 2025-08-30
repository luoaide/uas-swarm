package pilot

type Node struct {
	Role      Role
	Directive [5]Directive // A prioritized list of Directives.
}

func NewNode() *Node {
	return &Node{}
}

func (pilot *Node) AddDirective() {

}

func (pilot *Node) ReTask() {

}

// This file deals with high-level pilot functions.
