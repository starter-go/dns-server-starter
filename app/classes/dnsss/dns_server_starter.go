package dnsss

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/miekg/dns"
	"github.com/starter-go/application"
	"github.com/starter-go/dns-server-starter/app/utils"
	"github.com/starter-go/dns-server-starter/dnsss"
	"github.com/starter-go/vlog"
)

// ServerStarter 是 DNS 的启动器
type ServerStarter struct {

	//starter:component

	_as func(io.Closer) //starter:as(".")

	UDPPort int    //starter:inject("${dns.udp.port}")
	UDPHost string //starter:inject("${dns.udp.host}")

	Resolvers []dnsss.ResolverRegistry //starter:inject(".")

	timeout        time.Duration
	currentRuntime *innerDNSServerRuntime
}

func (inst *ServerStarter) _impl() (io.Closer, application.Lifecycle) {
	return inst, inst
}

func (inst *ServerStarter) Close() error {
	return inst.stop()
}

func (inst *ServerStarter) Life() *application.Life {
	l := &application.Life{
		OnCreate: inst.create,
		OnStart:  inst.start,
		OnStop:   inst.stop,
		// OnLoop:   inst.loop,
	}
	return l
}

func (inst *ServerStarter) create() error {
	inst.timeout = time.Second * 30
	return nil
}

func (inst *ServerStarter) start() error {
	vlog.Info("dnsss.ServerStarter.start()")
	chain := inst.loadResolverChain()
	rt := &innerDNSServerRuntime{
		starter: inst,
		chain:   chain,
	}
	older := inst.currentRuntime
	inst.currentRuntime = rt
	inst.closeOlder(older)
	rt.start()
	return inst.waitForRuntimeStarted(rt, inst.timeout)
}

func (inst *ServerStarter) loadResolverChain() dnsss.ResolverChain {
	builder := dnsss.ResolverChainBuilder{}
	src := inst.Resolvers
	for _, r1 := range src {
		tmp := r1.ListResolverRegistrations()
		builder.Add(tmp...)
	}
	builder.LogItems()
	return builder.Build()
}

func (inst *ServerStarter) closeOlder(older *innerDNSServerRuntime) error {
	if older == nil {
		return nil
	}
	older.stop()
	return inst.waitForRuntimeStopped(older, inst.timeout)
}

func (inst *ServerStarter) waitForRuntimeStopped(rt *innerDNSServerRuntime, timeout time.Duration) error {
	if rt == nil {
		return nil
	}
	const step = time.Millisecond * 500
	for t := timeout; t > 0; t -= step {
		if !rt.starting {
			return fmt.Errorf("bad state for waiting stopped")
		}
		if rt.stopped {
			return nil
		}
		time.Sleep(step)
	}
	return fmt.Errorf("timeout")
}

func (inst *ServerStarter) waitForRuntimeStarted(rt *innerDNSServerRuntime, timeout time.Duration) error {
	if rt == nil {
		return nil
	}
	const step = time.Millisecond * 500
	for t := timeout; t > 0; t -= step {
		if (!rt.starting) || rt.stopped {
			return fmt.Errorf("bad state for waiting started")
		}
		if rt.started {
			return nil
		}
		time.Sleep(step)
	}
	return fmt.Errorf("timeout")
}

func (inst *ServerStarter) stop() error {
	vlog.Info("dnsss.ServerStarter.stop()")
	older := inst.currentRuntime
	inst.currentRuntime = nil
	return inst.closeOlder(older)
}

// func (inst *ServerStarter) loop() error {
// 	return nil
// }

////////////////////////////////////////////////////////////////////////////////

type innerDNSServerRuntime struct {
	starter *ServerStarter
	server  *dns.Server

	chain dnsss.ResolverChain

	starting bool
	started  bool
	stopping bool
	stopped  bool
}

func (inst *innerDNSServerRuntime) start() {
	inst.starting = true
	go inst.run1()
}

func (inst *innerDNSServerRuntime) stop() {
	inst.stopping = true
	go inst.Close()

	const timeout = time.Second * 15
	inst.starter.waitForRuntimeStopped(inst, timeout)
}

func (inst *innerDNSServerRuntime) Close() error {
	ser := inst.server
	if ser == nil {
		return nil
	}
	return ser.Shutdown()
}

func (inst *innerDNSServerRuntime) run1() {

	inst.starting = true

	defer func() {
		x := recover()
		utils.HandleErrorX(x)
	}()
	defer func() {
		inst.stopping = true
		inst.stopped = true
	}()

	err := inst.run2()
	utils.HandleError(err)
}

func (inst *innerDNSServerRuntime) run2() error {

	addr := inst.computeUDPAddr()
	server := &dns.Server{
		Addr: addr,
		Net:  "udp",
	}
	server.TsigSecret = map[string]string{
		"axfr.": "so6ZGir4GPAqINNh9U5c3A==",
	}
	inst.server = server

	go func() {
		err := server.ListenAndServe()
		utils.HandleError(err)
		inst.stopping = true
	}()

	dns.HandleFunc(".", inst.handleDNSRequest)

	vlog.Info("serve DNS named-server at [%s]", addr)
	inst.started = true

	return inst.run3()
}

func (inst *innerDNSServerRuntime) run3() error {
	const step = time.Second
	for {
		if inst.stopping {
			break
		}
		if inst.stopped {
			break
		}
		time.Sleep(step)
	}
	return nil
}

func (inst *innerDNSServerRuntime) computeUDPAddr() string {
	host := inst.starter.UDPHost
	port := inst.starter.UDPPort
	return host + ":" + strconv.Itoa(port)
}

func (inst *innerDNSServerRuntime) handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {

	m2 := new(dns.Msg)
	m2.SetReply(r)

	ctx := new(dnsss.Context)
	ctx.Request = r
	ctx.Writer = w
	ctx.Response = m2

	err := inst.chain.Resolve(ctx)
	if err != nil {
		vlog.Warn("%v", err.Error())
	}
}

func (inst *innerDNSServerRuntime) deprecated_handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {

	qlist := r.Question
	// var question *dns.Question
	for _, q := range qlist {
		vlog.Info("dns.rr = %v", q.String())
		// question = &q
	}

	m2 := new(dns.Msg)
	m2.SetReply(r)

	if r.IsTsig() != nil {
		if w.TsigStatus() == nil {
			// *Msg r has an TSIG record and it was validated
			m2.SetTsig("axfr.", dns.HmacSHA256, 300, time.Now().Unix())
		} else {
			// *Msg r has an TSIG records and it was not validated
		}
	}

	// rr := inst.makeMockRR(question)
	// if rr != nil {
	// 	m2.Answer = []dns.RR{rr}
	// }

	w.WriteMsg(m2)
}
