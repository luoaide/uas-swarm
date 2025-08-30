package swarm

import (
	"archangel/node"
	"fmt"
	"testing"
	"time"
)

// func do(conf Config) {
// 	client := client.NewClient(
// 		option.WithToken(
// 			conf.ENVIRONMENT_TOKEN,
// 		),
// 		option.WithHTTPClient(XXX)
// 		),
// 	)
// 	client.Entities.LongPollEntityEvents(
// 		context.TODO(),
// 		&lattice.EntityEventRequest{
// 			SessionToken: "sessionToken",
// 		},
// 	)
// }

func TestExecuteMission(t *testing.T) {
	// create Myself as a node (i know if I'm an XNode, CNode, OpNode)
	self := node.NewCNode()
	interval := time.NewTicker(time.Second)
	// the complete channel can manually override the end of mission
	complete := make(chan bool)

	go executeMission(interval, complete, self)
}

func TestMain(t *testing.T) {
	// This is the main entry point. I need to
	// Read in Configurations
	// Join a swarm
	// Start collecting internal data "Reports" on myself (report)
	// Start ingesting external reports from other nodes (reporter)
	// Start sending "Votes" (elector)
	configurations := readInConfigurations()
	fmt.Println(configurations)

	// create Myself as a node (i know if I'm an XNode, CNode, OpNode)
	self := node.NewCNode()

	//swarmConfig := self.Node.AwaitSwarm()

	interval := time.NewTicker(time.Second)
	// the complete channel can manually override the end of mission
	complete := make(chan bool)

	go executeMission(interval, complete, self)

	//if a task comes in to terminate then set complete <- true
}
