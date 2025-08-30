package node

import (
	communicator "archangel/communicator"
	elector "archangel/elector"
	pilot "archangel/pilot"
	reporter "archangel/reporter"
)

// This is an operator. He must publish himself to the Lattice Mesh so that he can be tasked to execute things

// This is also the broker that might receive tasks from the GCS, accept them, and translate them into actionable
// tasks for the swarm.

// So the GCS operates separately from this coordination layer. But the coordination layer understands the
// movements of the Control UXV and appropriately changes the state of the swarm in response.

// This would run on the laptop that the operator uses or a Soldier-Borne Compute Module.
// Direct connection between the GCS and this compute device (if not on the same device)

type OpNode struct {
	Node
	ControlRadio *communicator.Radio
	RobotRadio   *communicator.Radio
	Network      *communicator.Network
}

func NewOpNode() *OpNode {
	return &OpNode{
		Node: Node{
			entityID:         "111",
			operationalSwarm: "test-swarm",
			Reporter:         reporter.NewNode(),
			Elector:          elector.NewNode(),
			Pilot:            pilot.NewNode(),
		},
		ControlRadio: communicator.NewRadio(),
		RobotRadio:   communicator.NewRadio(),
		Network:      communicator.NewNetwork(),
	}
}
