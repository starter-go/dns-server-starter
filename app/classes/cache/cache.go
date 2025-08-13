package cache

import (
	"github.com/starter-go/dns-server-starter/app/classes/orders"
	"github.com/starter-go/dns-server-starter/dnsss"
)

type Cache struct {

	//starter:component

	_as func(dnsss.ResolverRegistry) //starter:as(".")

}

func (inst *Cache) _impl() (dnsss.ResolverRegistry, dnsss.Resolver) {
	return inst, inst
}

func (inst *Cache) ListResolverRegistrations() []*dnsss.ResolverRegistration {
	r1 := &dnsss.ResolverRegistration{
		Name:     "cache.Cache",
		Resolver: inst,
		Order:    orders.Cache,
		Enabled:  true,
	}
	return []*dnsss.ResolverRegistration{r1}
}

func (inst *Cache) Resolve(ctx *dnsss.Context, next dnsss.ResolverChain) error {

	// TODO ...

	return next.Resolve(ctx)
}
