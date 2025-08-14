package units4dnsss

import (
	"net"
	"testing"
	"time"
)

func TestBindUDP(t *testing.T) {
	binder := &myUDPBinder{}
	err := binder.run()
	if err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
	t.Log("done")
}

////////////////////////////////////////////////////

type myUDPBinder struct{}

func (inst *myUDPBinder) run() error {

	const network = "udp"

	addr, err := net.ResolveUDPAddr(network, ":53")
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP(network, addr)
	if err != nil {
		return err
	}

	defer conn.Close()
	time.Sleep(time.Second * 5)
	return nil
}
