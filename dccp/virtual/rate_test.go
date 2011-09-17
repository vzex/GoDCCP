// Copyright 2011 GoDCCP Authors. All rights reserved.
// Use of this source code is governed by a 
// license that can be found in the LICENSE file.

package virtual

import (
	"testing"
	"time"
	"github.com/petar/GoDCCP/dccp"
	"github.com/petar/GoDCCP/dccp/ccid3"
)

func TestDropRate(t *testing.T) {
	hca, hcb, _ := NewLine(10)
	ccid := ccid3.CCID3{}
	/* cc := */ dccp.NewConnClient("Client", hca, ccid.NewSender(), ccid.NewReceiver(), 0)
	/* cs := */ dccp.NewConnServer("Server", hcb, ccid.NewSender(), ccid.NewReceiver())
	time.Sleep(10e9)
}
