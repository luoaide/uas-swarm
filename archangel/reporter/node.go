package reporter

import (
	"archangel/communicator"
)

// What is a Node from the Reporter's perspective?
//GPS
//Altimeter
//Camera?
//Proximity Sensor?
//Light Sensor?
//Mavlink reports????
//direction
//velocity
//INTENT

type Report struct {
}

type Peer struct {
	entityID string
	Report   Report
}

// do I need this?
type Updater interface {
	UpdateReport()
}

type Node struct {
	Semaphore      chan bool
	InternalStatus Report
	ExternalStatus []Report //make this a list of peers
}

// this is a constructor to build the empty Node struct back in the specific node (CNode, XNode, etc...) in the node package
// I'm effectively updating this object with the methods below
func NewNode() *Node {
	return &Node{}
}

func (node *Node) CollectInternal() {
	//generateInternalStatus(node)
}

// ExternalStatus needs setters/getters since we can have concurrency issues
// from multiple network connections
func (node *Node) addPeerReport() {
	//block
	//add
	//unblock
}

func (node *Node) CollectExternal(socket communicator.Listener) {
	// keep an open connection to the socket and accept one report at a time... then deconstruct it
	// build the golang object from reports.go and add/update ExternalStatus
	//incoming := socket.processReport()
	//node.addPeerReport(incoming)
}
