package dnsss

import (
	"fmt"
	"sort"
	"strings"

	"github.com/starter-go/vlog"
)

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

func (inst *ResolverChainBuilder) doSort(reverse bool) {
	var i sort.Interface = inst
	if reverse {
		i = sort.Reverse(i)
	}
	sort.Sort(i)
}

func (inst *ResolverChainBuilder) sort() {
	inst.doSort(false)
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

func (inst *ResolverChainBuilder) initWithItems(src []*ResolverRegistration) {
	dst := make([]*ResolverRegistration, 0)
	dst = append(dst, src...)
	inst.items = dst
}

func (inst *ResolverChainBuilder) LogItems() {

	sb := &strings.Builder{}

	tmp := new(ResolverChainBuilder)
	tmp.initWithItems(inst.items)
	tmp.doSort(true)
	list := tmp.items

	sb.WriteString("dnsss.ResolverChain:\n\n")
	for idx, it := range list {
		if !it.Enabled {
			continue
		}
		str := fmt.Sprintf("[Resolver index:%v order:%v name:'%v']\n", idx, it.Order, it.Name)
		sb.WriteString(str)
	}
	vlog.Info("%v", sb.String())
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
