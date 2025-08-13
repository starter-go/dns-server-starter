package monitor

import (
	"math"
	"strconv"
	"strings"

	"github.com/miekg/dns"
	"github.com/starter-go/dns-server-starter/app/classes/orders"
	"github.com/starter-go/dns-server-starter/dnsss"
	"github.com/starter-go/vlog"
)

type Monitor struct {

	//starter:component

	_as func(dnsss.ResolverRegistry) //starter:as(".")

	countRequest       int
	countResponse      int
	countAnswer        int
	countQuestion      int
	countResolving     int
	countResolvingSync int

	enableLogDetail bool
}

func (inst *Monitor) _impl() (dnsss.ResolverRegistry, dnsss.Resolver) {
	return inst, inst
}

func (inst *Monitor) ListResolverRegistrations() []*dnsss.ResolverRegistration {
	r1 := &dnsss.ResolverRegistration{
		Name:     "Monitor",
		Resolver: inst,
		Order:    orders.Monitor,
		Enabled:  true,
	}
	return []*dnsss.ResolverRegistration{r1}
}

func (inst *Monitor) Resolve(ctx *dnsss.Context, next dnsss.ResolverChain) error {
	inst.logRequest(ctx)
	err := next.Resolve(ctx)
	if err != nil {
		return err
	}
	inst.logResponse(ctx)
	inst.updateCountResolvingSync()
	return nil
}

func (inst *Monitor) logRequest(ctx *dnsss.Context) {
	req := ctx.Request
	list := req.Question
	if inst.enableLogDetail {
		for _, q := range list {
			str := inst.stringifyQuestion(&q)
			vlog.Info("question: %v", str)
		}
	}
	inst.countQuestion += len(list)
	inst.countRequest++
}

func (inst *Monitor) logResponse(ctx *dnsss.Context) {
	resp := ctx.Response
	list := resp.Answer
	if inst.enableLogDetail {
		for _, rr := range list {
			str := inst.stringifyRR(rr)
			vlog.Info("answer: %v", str)
		}
	}
	inst.countAnswer += len(list)
	inst.countResponse++
}

func (inst *Monitor) updateCountResolvingSync() {

	inst.countResolving++

	c1 := inst.countResolving
	c2 := inst.countResolvingSync
	gap := 100.0
	diff := math.Abs(float64(c1 - c2))
	if c1 > 5 {
		if diff < gap {
			return
		}
	}
	inst.countResolvingSync = c1

	b := strings.Builder{}
	b.WriteString("[count")

	b.WriteString(" req:")
	b.WriteString(strconv.Itoa(inst.countRequest))
	b.WriteString(" resp:")
	b.WriteString(strconv.Itoa(inst.countResponse))
	b.WriteString(" question:")
	b.WriteString(strconv.Itoa(inst.countQuestion))
	b.WriteString(" answer:")
	b.WriteString(strconv.Itoa(inst.countAnswer))

	b.WriteString("]")
	vlog.Info(b.String())
}

func (inst *Monitor) stringifyRR(rr dns.RR) string {
	return rr.String()
}

func (inst *Monitor) stringifyQuestion(q *dns.Question) string {
	return q.String()
}
