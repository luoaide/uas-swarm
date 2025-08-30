package communicator

type Radio struct {
}

type Listener interface {
	Receive([]byte)
}

type Sender interface {
	Send([]byte) error
}

func NewRadio() *Radio {
	return &Radio{}
}

func (radio *Radio) Send([]byte) {

}

func (radio *Radio) Receive([]byte) {

}

//requires information about where the RF device is and how that RF device communicates...
// might be easiest to set up that RF device as a Mavlink Component so that I have a pre-defined way of talking to it.
