package core

//var _ Conn = (*InArrayPort)(nil)

type InArrayPort struct {
	network *Network

	portName string
	fullName string
	array    []*Connection
	closed   bool
}

func (c *InArrayPort) isDrained() bool {
	for _, v := range c.array {
		if !v.isDrained() {
			return false
		}
	}
	return true
}

func (c *InArrayPort) IsEmpty() bool {
	return false
}

func (c *InArrayPort) receive(p *Process) *Packet {
	return nil
}

func (c *InArrayPort) IsClosed() bool {
	return c.closed
}

func (c *InArrayPort) resetForNextExecution() {}

func (c *InArrayPort) GetType() string {
	return "InArrayPort"
}

func (c *InArrayPort) GetArrayItem(i int) *Connection {
	if i >= len(c.array) {
		return nil
	}
	return c.array[i]
}

func (c *InArrayPort) SetArrayItem(c2 *Connection, i int) {
	if i >= len(c.array) {
		// add to .array to fit c2
		increaseBy := make([]*Connection, i-len(c.array)+1)
		c.array = append(c.array, increaseBy...)
	}
	c.array[i] = c2
}

func (c *InArrayPort) ArrayLength() int {
	return len(c.array)
}