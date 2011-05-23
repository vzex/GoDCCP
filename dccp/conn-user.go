// Copyright 2010 GoDCCP Authors. All rights reserved.
// Use of this source code is governed by a 
// license that can be found in the LICENSE file.

package dccp

import (
	"os"
)

// MaxBlockLen() returns the maximum size of a block that can be passed to WriteBlock
func (c *Conn) MaxBlockLen() int {
	return c.hc.MaxFootprint() - MAX_OPTIONS_SIZE - getFixedHeaderSize(DataAck, true) 
}

// WriteBlock blocks until the slice b is sent.
func (c *Conn) WriteBlock(b []byte) os.Error {
	if len(b) > c.MaxBlockLen() {
		return ErrTooBig
	}
	c.Lock()
	state := c.socket.GetState()
	c.Unlock()
	if state != OPEN {
		return os.EBADF
	}
	if len(b) == 0 {
		return nil
	}
	c.writeData <- b
	return nil
}

// ReadBlock blocks until the next packet of application data is received.
// It returns a non-nil error only if the connection has been closed.
func (c *Conn) ReadBlock() (b []byte, err os.Error) {
	b, ok := <-c.readApp
	if !ok {
		// The connection has been closed
		return nil, os.EBADF
	}
	return b, nil
}

// Close closes the connection, Section 8.3
func (c *Conn) Close() os.Error {
	c.Lock()
	state := c.socket.GetState()
	c.Unlock()
	if state == CLOSED || state == CLOSEREQ || state == CLOSING || state == TIMEWAIT {
		return nil
	}
	if state != OPEN {
		return os.EBADF
	}
	// Transition to CLOSING
	c.inject(c.generateClose())
	c.Lock()
	c.gotoCLOSING()
	c.Unlock()
	return nil
}
