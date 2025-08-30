package node

import (
	communicator "archangel/communicator"
	elector "archangel/elector"
	pilot "archangel/pilot"
	reporter "archangel/reporter"
)

type XNode struct {
	Node
	RobotRadio *communicator.Radio
}

func NewXNode() *XNode {
	return &XNode{
		Node: Node{
			entityID:         "111",
			operationalSwarm: "test-swarm",
			Reporter:         reporter.NewNode(),
			Elector:          elector.NewNode(),
			Pilot:            pilot.NewNode(),
		},
		RobotRadio: communicator.NewRadio(),
	}
}
