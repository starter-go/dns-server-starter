package cache

import (
	"github.com/starter-go/dns-server-starter/app/classes/orders"
	"github.com/starter-go/dns-server-starter/dnsss"
	"github.com/starter-go/vlog"
)

type Cache struct {

	//starter:component

	_as func(dnsss.ResolverRegistry) //starter:as(".")

	Enabled bool //starter:inject("${dnsss-resolver.cache.enabled}")

	// TTL in seconds

	MaxTTL int //starter:inject("${dnsss-resolver.cache.max-ttl}")
	MinTTL int //starter:inject("${dnsss-resolver.cache.min-ttl}")

	todoLogged bool
}

func (inst *Cache) _impl() (dnsss.ResolverRegistry, dnsss.Resolver) {
	return inst, inst
}

func (inst *Cache) getRegistration() *dnsss.ResolverRegistration {
	r1 := &dnsss.ResolverRegistration{
		Name:     "dnsss.cache",
		Resolver: inst,
		Order:    orders.Cache,
		Enabled:  inst.Enabled,
	}
	return r1
}

func (inst *Cache) ListResolverRegistrations() []*dnsss.ResolverRegistration {
	r1 := inst.getRegistration()

	if !inst.todoLogged {
		inst.todoLogged = true
		vlog.Warn("TODO: this component is NOT impl: " + r1.Name)
	}

	return []*dnsss.ResolverRegistration{r1}
}

func (inst *Cache) Resolve(ctx *dnsss.Context, next dnsss.ResolverChain) error {

	// TODO ...

	return next.Resolve(ctx)
}
