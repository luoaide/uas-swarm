package node

import (
	communicator "archangel/communicator"
	elector "archangel/elector"
	pilot "archangel/pilot"
	reporter "archangel/reporter"
	"fmt"
	"time"
)

type CNode struct {
	Node
	ControlRadio *communicator.Radio
	RobotRadio   *communicator.Radio
}

func NewCNode() *CNode {
	return &CNode{
		Node: Node{
			//I can write code later to handle these strings as "opt" paramaters
			entityID:         "111",
			operationalSwarm: "test-swarm",
			Reporter:         reporter.NewNode(),
			Elector:          elector.NewNode(),
			Pilot:            pilot.NewNode(),
		},
		ControlRadio: communicator.NewRadio(),
		RobotRadio:   communicator.NewRadio(),
	}
}

// func (node *CNode) Report() reporter.Report {
// 	//return node.Reporter()
// }

// To control concurrency: As a node, I have several well-defined data stores. These are:
// 		OnBoard Status (updated by reporter)
// 		Knowledge of Topology (updated by reporter)
// 		Orders (used by Pilot)
// 		Votes (owned by Elector)
// The only "hard time" I have is my window to send a report (Elector) to the network. If I'm not ready, then I miss my window
// and I will block updates until I make the next window.
// --- I will deconflict the writing and reading of these data stores using semaphores.

// --- means there is a goroutine WRITING to that datastore
// XXX means another goroutine needs to READ from that datastore

// NODE 1
// Reporter.Internal	|---|XXX|-|XXX|
// Reporter.External	|---|XXX|-|XXX|
// Pilot.Orders			|XXX|XXX|X|---|
// Elector.Ballot		|   |---|X|   |
// TIME					A   B   C~D   E

// NODE 2
// Reporter.Internal	X||---|XXX|-|XX
// Reporter.External	X||---|XXX|-|XX
// Pilot.Orders			-||XXX|XXX|X|--
// Elector.Ballot		 ||   |---|X|
// TIME					 EA   B   C~D

func (node CNode) Initialize() {
	// compute first mission objective
	// start collecting data and reports
	// this drives the default behavior. As things change (receive reports, vote, those methods are
	// responsible for blocking or change the default condition)
	go node.Pilot.ReTask()
	go node.Reporter.CollectInternal()
	go node.Reporter.CollectExternal(node.ControlRadio)
	go node.Reporter.CollectExternal(node.RobotRadio)
}

func (node CNode) Vote(t *time.Time) {
	fmt.Println(node.entityID, "Voting")
	// //STOP WRITING TO THIS NODE'S STATE STORES
	// reporterSemaphore := make(chan int, 1)
	// pilotSemaphore := make(chan int, 1)

	// //block all
	// reporterSemaphore <- 1

	// node.Elector.ComputeReport()

	// //resume all
	// <-reporterSemaphore

	// go node.ControlRadio.SendReport()
	// go node.RobotRadio.SendReport()
}

func (node *CNode) Retask() {
	//block reading of T&P
	node.Pilot.ReTask()
	//unblock reading of T&P
	//go node.Pilot.Execute()
}
