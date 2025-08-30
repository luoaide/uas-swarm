package node

import (
	elector "archangel/elector"
	pilot "archangel/pilot"
	reporter "archangel/reporter"
	"time"
)

type Node struct {
	entityID         string
	operationalSwarm string

	Reporter *reporter.Node
	Elector  *elector.Node
	Pilot    *pilot.Node
}

type Config struct {
	MyOffset time.Duration
}

func (n *Node) AwaitSwarm() Config {
	return Config{
		MyOffset: 45,
	}
}

// A Voter is any node in the swarm that can send a Report/Ballot to the warm.
// This means that the node can report, vote and execute orders
// Each unique type of node must be able to execute an epoch using the radios, sensors
// and autopilot available onboard.
type Voter interface {
	Vote(t *time.Time)
}

// a Taskable Node is one that can be retasked by the consensus reached in the swarm. Retasking is completed
// after each epoch based on the votes received and a node's status in the network.
type Taskable interface {
	ReTask()
}
