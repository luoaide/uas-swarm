package elector

// an Elector takes one epoch of reports and nominations and outputs a single vote

// There is some cross-over between reporter functionality and elector functionality. Who should iterate
// across all received reports to generate new directives and commands to pass to the pilot?

// TEST NOMINATION

// testNomination := Nomination{
// 	nominee	"1234"  //entityID of the node nominated
// 	command {

// 	}
// }

type Nomination struct {
	//votes from others (ensure there is a quorum voting on this nomination before I decide to implement or not)
}

// ElectorNode
type Node struct {
	//My Role
	//Votes?
}

func NewNode() *Node {
	return &Node{}
}
