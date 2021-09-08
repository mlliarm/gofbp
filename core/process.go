package core

//import (
//	"sync"
//)

type Process struct {
	Name string
	//procs   map[string]Process
	Network    *Network
	inPorts    map[string]*InPort
	outPorts   map[string]*OutPort
	logFile    string
	component  Component
	ownedPkts  int
	done       bool
	allDrained bool
	hasData    bool
}

func (p *Process) OpenInPort(s string) *InPort {
	return p.inPorts[s]
}

func (p *Process) OpenOutPort(s string) *OutPort {
	return p.outPorts[s]
}

func (p *Process) Run(net *Network) {

	p.component.OpenPorts(p)

	for !p.done {
		p.hasData = false
		p.allDrained = true
		for _, v := range p.inPorts {
			v.Conn.mtx.Lock()
			if !v.Conn.IsEmpty() {
				p.hasData = true
			}
			if !v.Conn.closed {
				p.allDrained = false
			}
			v.Conn.mtx.Unlock()
		}

		if len(p.inPorts) == 0 || !p.allDrained {

			p.component.Execute(p) // activate component Execute logic

			if p.ownedPkts > 0 {
				panic(p.Name + "deactivated without disposing of all owned packets")
			}
		}
		if p.allDrained {
			break
		}
	}
	p.done = true
	for _, v := range p.inPorts {
		v.Conn.mtx.Lock()
		if !(v.Conn.closed && !v.Conn.IsEmpty()) {
			p.done = false
		}
		v.Conn.mtx.Unlock()
	}

	for _, v := range p.outPorts {
		v.Conn.mtx.Lock()
		v.Conn.UpStrmCnt--
		if v.Conn.UpStrmCnt == 0 {
			v.Conn.closed = true
		}
		v.Conn.mtx.Unlock()
	}

}

func (p *Process) Create(s string) *Packet {
	var pkt *Packet = new(Packet)
	pkt.Contents = s
	pkt.owner = p
	p.ownedPkts++
	return pkt
}

func (p *Process) Discard(pkt *Packet) {
	p.ownedPkts--
}
