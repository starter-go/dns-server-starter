package dnsss

type Resolver interface {
	Resolve(ctx *Context, next ResolverChain) error
}

type ResolverRegistration struct {
	Name     string
	Resolver Resolver
	Enabled  bool
	Order    int
}

type ResolverRegistry interface {
	ListResolverRegistrations() []*ResolverRegistration
}
