package POC

type KeySet struct {

}

type State struct {

}

type Emulator struct {
	curFrame int
	framePerSec int
}

func (e Emulator) runFrame(k KeySet) {

}

func (e Emulator) loadState(s State) {

}

func (e Emulator) saveState() State {
	var a = State{}
	return a
}


// ---------------------------------------------
// Message definition
const ( // iota is reset to 0
	msg_start = iota
	msg_keyframe = iota
	msg_resync = iota
)

type Message struct {
	messageType int
	data Data
}

type Data struct {
	keyset *KeySet
}

func (data Data) getInt() int {
	return 0
}

func (data Data) getKeySet() *KeySet {
	return data.keyset
}





// ----------------------------------------------

type Time float32

func (t Time) toFloat() float32 {
	return float32(t)
}

type Client struct {
	emulator Emulator
	lastRunTime Time
	id int
}

func (client Client) init() {

}

func (client Client) update(dt Time) {

}

func (client Client) dispatch() {

}

// -------------------------------------------------

type ClientRep struct {
	id int
}

type ReplayFrame struct {
	data []KeySet
}

type ReplayData struct {
	frameCount int
	frames []ReplayFrame
}

type Server struct {
	clients []ClientRep
	data ReplayData
}

func (server Server) update(t Time) {
	// call dispatch
}

func (server Server) processMessage(client ClientRep, msg Message) {
	// 
}

func (server Server) dispatch(msgs []Message) {
	// call system to dispatch msgs.
}

func (server Server) receive() {
	// get messages from System and process them
}

func (server Server) addClient(client ClientRep) {

}

func (server Server) removeClient(client ClientRep) {

}

// -----------------------------------------------------

type System struct {
	demo *Demo
	clients map[int]ClientRep
}

func (system System) serverSend(client ClientRep, msgs []Message) {

}

func (system System) serverRecv(client ClientRep) []Message {
	return nil
}

func (system System) clientSend(id int, msgs []Message) {

}

func (system System) clientRecv(id int) []Message {
	return nil
}

var sys System

// -------------------------------------------------------

type Demo struct {
	clients []Client
}

func (demo Demo) ss() {}

func (demo Demo) run(clientCount int) {
	sys = System{demo: &demo}
	for i := 0; i < clientCount; i++ {
		demo.clients = append(demo.clients, Client{id: i})
		sys.clients[i] = ClientRep{id: i}
	}
}
