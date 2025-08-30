package pilot

// This file deals with the actions that the pilot implements

// I should operate based on "Roles". these are essentially wrappers around Mavlink commands so that "Votes"
// sent between nodes can strip out some of the unneccessary coordination... an leave implementation of a role
// to the onboard computer for an executer.
// Roles:
// 1. Human to Ingest Comms Relay Node
// 2. Vehicle Comms Network
// 3. Strike Platform (OWA)
// 4. Strike Platform (Dropper)
// 5. Intercept (block incoming threats)
// 6. Relay Node
// 7. Observer (tied to another agent)
// 8. Swarm Eyes (the virtual agent that the operator technically pilots and "leads" the rest of the pack to the Area of Interest).

// Nominations (each one needs a format):
// Types:  direct (hey you, do this)
//         indirect (someone do this)
// 1. Relieve me (take over my role)
// 2. Relieve X
// 2. Observe Me
// 3. Observe X
// 3. Move out of my way (Change position of Follow Agent)
// 4. Assume X Role
// 5.

type Role struct {
}

type Executer interface {
	Execute()
}

type Command struct {
}

type Directive struct {
	Command Command
}

func (directive *Directive) Execute() {
	// implement the directive.
}
