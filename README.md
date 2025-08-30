# Archangel
Open Source swarm protocol for multi-UXV coordination.

# Strategy
/swarm is the service that runs w/ the operator to serve as the ingress point to the swarm
/cnodes have a data link to the operator
/xnodes execute tasks and communicate among each other and to cnodes but not to the operator

Guide       ->      Anduril Lattice / GCS
Coordinate  ->      Archangel (This layer)
Terminal    ->      Peregrin (AI2C)
Commands    ->      Mavlink
FC          ->      PX4/Betaflight

Each xnode and cnode has:
1. Pilot: converts decisions made by Archangel to flight instructions (needs Mavlink)
2. Network
3. Elector
4. Pundit
5. Reporter

Each Swarm has:
1. One "Taskable" Swarm Entity in Lattice
2. Multiple non-taskable entities associated with the Swarm


# NATO OPERATIONS RESEARCH PAPER
Abstract (150 to 250 words)

The span of control of individual drone operators currently limits the quantity of unmanned systems that can be employed in a tactical scenario. Operators can achieve more effects and more efficient effects by increasing the quantity of systems under their control and broadening their focus. By removing the task of piloting a vehicle, operators can focus on desired effects and better manage their resources. To realize these efficiencies, unmanned vehicles need a framework through which to autonomously communicate and coordinate their interactions. This paper introduces a voting heuristic for multi-agent unmanned vehicle interactions that enables human-in-the-loop decision-making for autonomous missions. The protocol described in this paper works in sub-second "election cycles." In each election cycle, every active node (unmanned vehicle) reports on its current condition, votes on proposals made during the previous election cycle, and makes proposals to be adjudicated in the next. This mechanism provides a best-effort immutable state for the entire network, embraces imperfect information, uses consensus to inform actions, and is modular enough to integrate other autonomy solutions. The implementation of this protocol discussed in this paper consists of a network controller to handle radio-frequency (RF) traffic between the nodes in the network, a flight stack controller to issue commands to the onboard flight controller, an Elector to make voting decisions, a Pundit to predict elections and optimize positioning, and a Reporter that ingests environmental and node data to generate reports that are shared with the network to inform elections.

Each election period, every alive node votes on the nominations made during the previous election cycle. Nominations are either made by an operator (prioritized) or by any vehicle in the system. Election cycles arbitrarily happen every 1 second. 

If a node loses connection with the rest of the network by not participating in a series of elections, it and the rest of the network are aware of this based on its non-participation.

Thoughts:
1. As a node, when do I decide to move into the next election cycle? Possible options:
    Time interval
    When I have some kind of clarity
    when I reach a "quorom"

If I'm a leader node, I can send the first report of an election cycle unprompted, or everyone can be trying to send that first report of an election cycle unless someone beats me to it. we can have multiple pockets consolidating reports and sending them onward.

Concepts:
1. report folding
2. companion computer (issues commands to the on-board flight stack)
3. Placement (in 3D space) -- relay + observer + attack angle (90degress for dropper)
4. Role
5. Zones of placement (holding, enroute, on strike, )

Roles:
1. Human to Ingest Comms Relay Node
2. Vehicle Comms Network
3. Strike Platform (OWA)
4. Strike Platform (Dropper)
5. Intercept (block incoming threats)
6. Relay Node
7. Observer

Software Components:
1. Comms Controller
    reporting
    receiving reports
2. Flight Stack Controller (issues commands to the local vehicle)
3. Report Generator (analyzing current node conditions and generates status reports)
4. Pilot (based on role, and last election, actually sends PX4 commands to the flight controller... needs a plugg-able VNAV component)
5. Pundit (predicts the outcome of the next few elections and pre-positions to optimize)


## Design
This is a control layer utility acting as an independant control mechanism for a large swarm of autonomous robots.
1. Each robot knows only its surrounding environment. Based on:
    1. On-board sensors
    2. Link data about the FANET
2. Each robot is likely controled w/ PX4 flight stack. This "brain" software runs on a companion computer and provides commands via Mavlink to the FC.
3. Any drone can make a nomination
4. status report from each drone is an immutable structure and contains STATUS, NOMINATOINS, VOTES

Graph theory mixed w/ comms

a virtual navigator provides mean influence in terms of velocity and acceleration
obstacle avoidance is the responsibility of overall agents
acceleration to avoid obstacles may result in a need for adjacent agents to move

Do we define roles in the swarm? like a fighter jet formation when there is a leader and then basically a wall of strike options for kinetic attacks... but its all self-healing

what is the actual military utility?

in an attack formation, X systems approach the target from a variety of angles, then at the last moment a final attack decision is made.

buf = buffer space required for safe operation around each robot


Px(t)
Py(t)

## Definitions

Must:
1. Run on an on-board Raspberry Pi (prefer zero)
2. Connect RF transmitter/received to Controller instead of the Flight Controller.
3. Connect Controller to the FC.

Roles:
1. Human to Ingest Comms Relay Node
2. Vehicle Comms Network
3. Strike Platform (OWA)
4. Strike Platform (Dropper)
5. Intercept (block incoming threats)
6. Relay Node
7. Observer (tied to another agent)
8. Swarm Eyes (the virtual agent that the operator technically pilots and "leads" the rest of the pack to the Area of Interest).

Nominations (each one needs a format):
Types:  direct (hey you, do this)
        indirect (someone do this)
1. Relieve me (take over my role)
2. Relieve X
2. Observe Me
3. Observe X
3. Move out of my way (Change position of Follow Agent)
4. Assume X Role
5. 


## Resources
1. Enhanced multi agent coordination algorithm for drone swarm patrolling in durian orchards