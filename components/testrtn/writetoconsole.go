package testrtn

import (
	"fmt"

	"github.com/jpaulm/gofbp/core"
)

type WriteToConsole struct {
	ipt core.InputConn
	out core.OutputConn
}

func (writeToConsole *WriteToConsole) Setup(p *core.Process) {
	writeToConsole.ipt = p.OpenInPort("IN")
	writeToConsole.out = p.OpenOutPortOptional("OUT")
}

func (WriteToConsole) MustRun() {}

func (writeToConsole *WriteToConsole) Execute(p *core.Process) {

	for {
		var pkt = p.Receive(writeToConsole.ipt)
		if pkt == nil {
			break
		}
		if pkt.PktType == core.OpenBracket {
			fmt.Println("Open", pkt.Contents)
		} else {
			if pkt.PktType == core.CloseBracket {
				fmt.Println("Close", pkt.Contents)
			} else {
				fmt.Println(pkt.Contents)
			}
		}

		if writeToConsole.out.IsConnected() {
			p.Send(writeToConsole.out, pkt)
		} else {
			p.Discard(pkt)
		}
	}

}
