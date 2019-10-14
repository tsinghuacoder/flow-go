package crypto

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// maps the interval [0..n-1] into [0..c-1,c+1..n] if c>1
// maps the interval [0..n-1] into [1..n] if c=0
func index(current int, loop int) int {
	if loop < current {
		return loop
	}
	return loop + 1
}

const (
	dkgType int = iota
	tsType
)

type toProcess struct {
	orig    int
	msgType int
	msg     interface{}
}

// This is a testing function
// it simulates sending a message from one node to another
func send(orig int, dest int, msgType int, msg interface{}, chans []chan *toProcess) {
	log.Infof("%d Sending to %d:\n", orig, dest)
	log.Debug(msg)
	newMsg := &toProcess{orig, msgType, msg}
	chans[dest] <- newMsg
}

// This is a testing function
// it simulates broadcasting a message from one node to all nodes
func broadcast(orig int, msgType int, msg interface{}, chans []chan *toProcess) {
	log.Infof("%d Broadcasting:", orig)
	log.Debug(msg)
	newMsg := &toProcess{orig, msgType, msg}
	for i := 0; i < len(chans); i++ {
		if i != orig {
			chans[i] <- newMsg
		}
	}
}

// This is a testing function
// It simulates processing incoming messages by a node
func dkgProcessChan(current int, dkg []DKGstate, chans []chan *toProcess,
	quit chan int, t *testing.T) {
	for {
		select {
		case newMsg := <-chans[current]:
			log.Infof("%d Receiving from %d:", current, newMsg.orig)
			out := dkg[current].ReceiveDKGMsg(newMsg.orig, newMsg.msg.(DKGmsg))
			out.processDkgOutput(current, dkg, chans, t)
		// if timeout, stop and finalize
		case <-time.After(time.Second):
			log.Infof("%d quit \n", current)
			dkg[current].EndDKG()
			quit <- 1
			return
		}
	}
}

// This is a testing function
// It processes the output of a the DKG library
func (out *DKGoutput) processDkgOutput(current int, dkg []DKGstate,
	chans []chan *toProcess, t *testing.T) {
	assert.Nil(t, out.err)
	if out.err != nil {
		log.Error("DKG output error: " + out.err.Error())
	}

	assert.Equal(t, out.result, valid, "they should be equal")

	for _, msg := range out.action {
		if msg.broadcast {
			broadcast(current, dkgType, msg.data, chans)
		} else {
			send(current, msg.dest, dkgType, msg.data, chans)
		}
	}
}

// Testing the happy path of Feldman VSS by simulating a network of n nodes
func TestFeldmanVSS(t *testing.T) {
	log.SetLevel(log.ErrorLevel)
	log.Debug("Feldman VSS starts")
	// number of nodes to test
	n := 5
	lead := 0
	dkg := make([]DKGstate, n)
	quit := make(chan int)
	chans := make([]chan *toProcess, n)

	// create DKG in all nodes
	for current := 0; current < n; current++ {
		var err error
		dkg[current], err = NewDKG(FeldmanVSS, n, current, lead)
		assert.Nil(t, err)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
	// create the node channels
	for i := 0; i < n; i++ {
		chans[i] = make(chan *toProcess, 10)
		go dkgProcessChan(i, dkg, chans, quit, t)
	}
	// start DKG in all nodes but the leader
	seed := []byte{1, 2, 3}
	for current := 0; current < n; current++ {
		if current != lead {
			out := dkg[current].StartDKG(seed)
			out.processDkgOutput(current, dkg, chans, t)
		}
	}
	// start the leader (this avoids a data racing issue)
	out := dkg[lead].StartDKG(seed)
	out.processDkgOutput(lead, dkg, chans, t)

	// TODO: check all public keys are consistent
	// TODO: check the reconstructed key is equal to a_0

	// this loop synchronizes the main thread to end all DKGs
	for i := 0; i < n; i++ {
		<-quit
	}
}
