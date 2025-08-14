package responder

import (
	"fmt"

	"github.com/starter-go/dns-server-starter/app/classes/orders"
	"github.com/starter-go/dns-server-starter/dnsss"
)

type DnsssResponder struct {

	//starter:component

	_as func(dnsss.ResolverRegistry) //starter:as(".")

}

func (inst *DnsssResponder) _impl() (dnsss.ResolverRegistry, dnsss.Resolver) {
	return inst, inst
}

func (inst *DnsssResponder) ListResolverRegistrations() []*dnsss.ResolverRegistration {

	r1 := &dnsss.ResolverRegistration{
		Name:     "DnsssResponder",
		Resolver: inst,
		Order:    orders.Responder,
		Enabled:  true,
	}

	return []*dnsss.ResolverRegistration{r1}
}

func (inst *DnsssResponder) Resolve(ctx *dnsss.Context, next dnsss.ResolverChain) error {

	err := next.Resolve(ctx)
	if err != nil {
		return err
	}

	if ctx.HasAnswer() {
		msg := ctx.Response
		return ctx.Writer.WriteMsg(msg)

	}

	return fmt.Errorf("no response answer")
}
