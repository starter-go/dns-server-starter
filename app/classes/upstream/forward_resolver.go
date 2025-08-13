package upstream

import (
	"bytes"
	"fmt"
	"strings"
	"sync"

	"github.com/miekg/dns"
	"github.com/starter-go/dns-server-starter/app/classes/orders"
	"github.com/starter-go/dns-server-starter/dnsss"
)

type ForwardResolver struct {

	//starter:component

	_as func(dnsss.ResolverRegistry) //starter:as(".")

	// 以逗号分隔的 DNS 服务器地址， like 'host:port', 其中的 port 是可选的， 默认值为 53
	DnsUpstreamServers string //starter:inject("${dns.resolver.upstream.servers}")

	frc      *myForwardResolverContext
	frcMutex sync.Mutex
}

func (inst *ForwardResolver) _impl() (dnsss.ResolverRegistry, dnsss.Resolver) {
	return inst, inst
}

func (inst *ForwardResolver) ListResolverRegistrations() []*dnsss.ResolverRegistration {
	r1 := &dnsss.ResolverRegistration{
		Name:     "upstream.ForwardResolver",
		Resolver: inst,
		Order:    orders.Upstream,
		Enabled:  true,
	}
	return []*dnsss.ResolverRegistration{r1}
}

func (inst *ForwardResolver) parseServerAddressList(str string) ([]string, error) {

	src := strings.Split(str, ",")
	dst := make([]string, 0)

	for _, item1 := range src {
		item1 = strings.TrimSpace(item1)
		if item1 == "" {
			continue
		}
		item2, err := inst.parseServerAddressItem(item1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}

	return dst, nil
}

func (inst *ForwardResolver) parseServerAddressItem(str string) (string, error) {

	bin := []byte(str)
	idx := bytes.IndexByte(bin, ':')

	if idx < 0 {
		want := " want: 'ip.addr:port'"
		have := " have: '" + str + "'"
		return "", fmt.Errorf("bad format of DNS-server-addr," + want + "," + have)
	}

	str1 := string(bin[0:idx])
	str2 := string(bin[idx+1:])

	if str1 == "" {
		str1 = "127.0.0.1" // default host
	}
	if str2 == "" {
		str2 = "53" // default port
	}

	return str1 + ":" + str2, nil
}

func (inst *ForwardResolver) loadFRContext() (*myForwardResolverContext, error) {

	mtx := &inst.frcMutex
	mtx.Lock()
	defer mtx.Unlock()

	ctx := inst.frc
	if ctx != nil {
		return ctx, nil
	}

	str := inst.DnsUpstreamServers
	addrlist, err := inst.parseServerAddressList(str)
	if err != nil {
		return nil, err
	}

	ctx = new(myForwardResolverContext)
	client := new(dns.Client)
	ctx.client = client
	ctx.servers = addrlist

	inst.frc = ctx
	return ctx, nil
}

func (inst *ForwardResolver) getFRContext() (*myForwardResolverContext, error) {
	frc := inst.frc
	if frc == nil {
		frc2, err := inst.loadFRContext()
		if err != nil {
			return nil, err
		}
		frc = frc2
		// inst.frc = frc2
	}
	return frc, nil
}

func (inst *ForwardResolver) Resolve(ctx *dnsss.Context, next dnsss.ResolverChain) error {

	frc, err := inst.getFRContext()
	if err != nil {
		return err
	}

	client := frc.client
	addrlist := frc.servers
	req := ctx.Request

	for _, addr := range addrlist {
		resp, _, err := client.Exchange(req, addr)
		if err == nil && resp != nil {
			ctx.Response = resp
			return nil
		}

		// else if err != nil {
		// 	vlog.Error("%v", err.Error())
		// }
	}

	return next.Resolve(ctx)
}

////////////////////////////////////////////////////////////////////////////////

type myForwardResolverContext struct {
	servers []string // items, like '1.2.3.40:53'
	client  *dns.Client
}
