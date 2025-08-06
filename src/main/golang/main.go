package main

import (
	"os"

	"github.com/starter-go/dns-server-starter/modules/dnsss"
	"github.com/starter-go/starter"
)

func main() {
	m := dnsss.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
