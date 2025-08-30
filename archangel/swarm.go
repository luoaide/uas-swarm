package swarm

import (
	"archangel/node"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type Config struct {
	LATTICE_ENDPOINT  string `yaml:"lattice-endpoint"`
	ENVIRONMENT_TOKEN string `yaml:"environment-token"`
	SANDBOXES_TOKEN   string `yaml:"sandboxes-token"`
	IDENTIFIER        string `yaml:"identifier"`
}

func check(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
}

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

func readInConfigurations() Config {
	conf := Config{}

	dat, err := os.ReadFile("./values.yaml")
	check(err)

	err = yaml.UnmarshalStrict([]byte(dat), &conf)
	check(err)
	fmt.Println(conf)

	return conf
}

// A Node is anyone in the swarm with a voting, reporting, and executing role
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
func executeMission(interval *time.Ticker, complete <-chan bool, self node.Voter) {
	for {
		select {
		case <-complete:
			return
		case t := <-interval.C:
			go self.Vote(&t)
			fmt.Println("Tick at", t)
		}
	}
}

func main() {
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

	swarmConfig := self.Node.AwaitSwarm()

	interval := time.NewTicker(swarmConfig.MyOffset)
	// the complete channel can manually override the end of mission
	complete := make(chan bool)

	go executeMission(interval, complete, self)

	//if a task comes in to terminate then set complete <- true
}
