package monitor

import (
	"github.com/starter-go/dns-server-starter/app/classes/orders"
	"github.com/starter-go/dns-server-starter/dnsss"
)

type ErrorFilter struct {

	//starter:component

	_as func(dnsss.ResolverRegistry) //starter:as(".")

	Enabled   bool //starter:inject("${dnsss-resolver.error-filter.enabled}")
	QuietMode bool //starter:inject("${dnsss-resolver.error-filter.quiet-mode}")
}

func (inst *ErrorFilter) _impl() (dnsss.ResolverRegistry, dnsss.Resolver) {
	return inst, inst
}

func (inst *ErrorFilter) ListResolverRegistrations() []*dnsss.ResolverRegistration {

	// inst.QuietMode = true

	r1 := &dnsss.ResolverRegistration{
		Name:     "dnsss.error-filter",
		Resolver: inst,
		Order:    orders.ErrorFilter,
		Enabled:  inst.Enabled,
	}
	return []*dnsss.ResolverRegistration{r1}
}

func (inst *ErrorFilter) Resolve(ctx *dnsss.Context, next dnsss.ResolverChain) error {

	err := next.Resolve(ctx)

	if err != nil {
		if inst.QuietMode {
			err = nil
		}
	}

	return err
}
