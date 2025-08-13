package dnsss

import "sort"

////////////////////////////////////////////////////////////////////////////////
// ResolverChain

type ResolverChain interface {
	Resolve(ctx *Context) error
}

////////////////////////////////////////////////////////////////////////////////
// ResolverChainBuilder

type ResolverChainBuilder struct {
	items []*ResolverRegistration
}

func (inst *ResolverChainBuilder) Add(list ...*ResolverRegistration) *ResolverChainBuilder {
	for _, item := range list {
		if inst.isItemReady(item) {
			inst.items = append(inst.items, item)
		}
	}
	return inst
}

func (inst *ResolverChainBuilder) isItemReady(item *ResolverRegistration) bool {
	if item == nil {
		return false
	}

	if !item.Enabled {
		return false
	}

	if item.Resolver == nil {
		return false
	}

	return true
}

func (inst *ResolverChainBuilder) sort() {
	sort.Sort(inst)
}
func (inst *ResolverChainBuilder) Len() int {
	return len(inst.items)
}
func (inst *ResolverChainBuilder) Less(i1, i2 int) bool {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	return o1.Order > o2.Order
}
func (inst *ResolverChainBuilder) Swap(i1, i2 int) {
	list := inst.items
	list[i1], list[i2] = list[i2], list[i1]
}

func (inst *ResolverChainBuilder) Build() ResolverChain {

	inst.sort()
	list := inst.items
	var chain ResolverChain
	chain = new(innerChainEnding)

	for _, item := range list {
		node := new(innerChainNode)
		node.next = chain
		node.reg = item
		node.resolver = item.Resolver
		chain = node
	}

	return chain
}

////////////////////////////////////////////////////////////////////////////////
// innerChainNode

type innerChainNode struct {
	reg      *ResolverRegistration
	resolver Resolver
	next     ResolverChain
}

func (inst *innerChainNode) Resolve(ctx *Context) error {
	return inst.resolver.Resolve(ctx, inst.next)
}

////////////////////////////////////////////////////////////////////////////////
// innerChainEnding

type innerChainEnding struct{}

func (inst *innerChainEnding) Resolve(ctx *Context) error {
	return nil // NOP
}

////////////////////////////////////////////////////////////////////////////////
// EOF
