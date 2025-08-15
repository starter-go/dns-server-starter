package main4dnsss
import (
    p9fd301bff "github.com/starter-go/dns-server-starter/app/classes/cache"
    pc47b01c43 "github.com/starter-go/dns-server-starter/app/classes/dnsss"
    p11d85ef56 "github.com/starter-go/dns-server-starter/app/classes/monitor"
    p2eec8228d "github.com/starter-go/dns-server-starter/app/classes/responder"
    peceb8058a "github.com/starter-go/dns-server-starter/app/classes/upstream"
    p9ee9f73b8 "github.com/starter-go/dns-server-starter/dnsss"
     "github.com/starter-go/application"
)

// type p9fd301bff.Cache in package:github.com/starter-go/dns-server-starter/app/classes/cache
//
// id:com-9fd301bff6e7228e-cache-Cache
// class:class-9ee9f73b8d049a5dcf2cc5fdc03de167-ResolverRegistry
// alias:
// scope:singleton
//
type p9fd301bff6_cache_Cache struct {
}

func (inst* p9fd301bff6_cache_Cache) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9fd301bff6e7228e-cache-Cache"
	r.Classes = "class-9ee9f73b8d049a5dcf2cc5fdc03de167-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9fd301bff6_cache_Cache) new() any {
    return &p9fd301bff.Cache{}
}

func (inst* p9fd301bff6_cache_Cache) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9fd301bff.Cache)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.MaxTTL = inst.getMaxTTL(ie)
    com.MinTTL = inst.getMinTTL(ie)


    return nil
}


func (inst*p9fd301bff6_cache_Cache) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${dnsss-resolver.cache.enabled}")
}


func (inst*p9fd301bff6_cache_Cache) getMaxTTL(ie application.InjectionExt)int{
    return ie.GetInt("${dnsss-resolver.cache.max-ttl}")
}


func (inst*p9fd301bff6_cache_Cache) getMinTTL(ie application.InjectionExt)int{
    return ie.GetInt("${dnsss-resolver.cache.min-ttl}")
}



// type pc47b01c43.ServerStarter in package:github.com/starter-go/dns-server-starter/app/classes/dnsss
//
// id:com-c47b01c433b94413-dnsss-ServerStarter
// class:class-f98ed07a4d5f50f7de1410d905f1477f-Closer
// alias:
// scope:singleton
//
type pc47b01c433_dnsss_ServerStarter struct {
}

func (inst* pc47b01c433_dnsss_ServerStarter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c47b01c433b94413-dnsss-ServerStarter"
	r.Classes = "class-f98ed07a4d5f50f7de1410d905f1477f-Closer"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc47b01c433_dnsss_ServerStarter) new() any {
    return &pc47b01c43.ServerStarter{}
}

func (inst* pc47b01c433_dnsss_ServerStarter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc47b01c43.ServerStarter)
	nop(ie, com)

	
    com.UDPPort = inst.getUDPPort(ie)
    com.UDPHost = inst.getUDPHost(ie)
    com.Resolvers = inst.getResolvers(ie)


    return nil
}


func (inst*pc47b01c433_dnsss_ServerStarter) getUDPPort(ie application.InjectionExt)int{
    return ie.GetInt("${dns.udp.port}")
}


func (inst*pc47b01c433_dnsss_ServerStarter) getUDPHost(ie application.InjectionExt)string{
    return ie.GetString("${dns.udp.host}")
}


func (inst*pc47b01c433_dnsss_ServerStarter) getResolvers(ie application.InjectionExt)[]p9ee9f73b8.ResolverRegistry{
    dst := make([]p9ee9f73b8.ResolverRegistry, 0)
    src := ie.ListComponents(".class-9ee9f73b8d049a5dcf2cc5fdc03de167-ResolverRegistry")
    for _, item1 := range src {
        item2 := item1.(p9ee9f73b8.ResolverRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p11d85ef56.ErrorFilter in package:github.com/starter-go/dns-server-starter/app/classes/monitor
//
// id:com-11d85ef56047f359-monitor-ErrorFilter
// class:class-9ee9f73b8d049a5dcf2cc5fdc03de167-ResolverRegistry
// alias:
// scope:singleton
//
type p11d85ef560_monitor_ErrorFilter struct {
}

func (inst* p11d85ef560_monitor_ErrorFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-11d85ef56047f359-monitor-ErrorFilter"
	r.Classes = "class-9ee9f73b8d049a5dcf2cc5fdc03de167-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p11d85ef560_monitor_ErrorFilter) new() any {
    return &p11d85ef56.ErrorFilter{}
}

func (inst* p11d85ef560_monitor_ErrorFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p11d85ef56.ErrorFilter)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.QuietMode = inst.getQuietMode(ie)


    return nil
}


func (inst*p11d85ef560_monitor_ErrorFilter) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${dnsss-resolver.error-filter.enabled}")
}


func (inst*p11d85ef560_monitor_ErrorFilter) getQuietMode(ie application.InjectionExt)bool{
    return ie.GetBool("${dnsss-resolver.error-filter.quiet-mode}")
}



// type p11d85ef56.Monitor in package:github.com/starter-go/dns-server-starter/app/classes/monitor
//
// id:com-11d85ef56047f359-monitor-Monitor
// class:class-9ee9f73b8d049a5dcf2cc5fdc03de167-ResolverRegistry
// alias:
// scope:singleton
//
type p11d85ef560_monitor_Monitor struct {
}

func (inst* p11d85ef560_monitor_Monitor) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-11d85ef56047f359-monitor-Monitor"
	r.Classes = "class-9ee9f73b8d049a5dcf2cc5fdc03de167-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p11d85ef560_monitor_Monitor) new() any {
    return &p11d85ef56.Monitor{}
}

func (inst* p11d85ef560_monitor_Monitor) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p11d85ef56.Monitor)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.EnableLogDetail = inst.getEnableLogDetail(ie)


    return nil
}


func (inst*p11d85ef560_monitor_Monitor) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${dnsss-resolver.monitor.enabled}")
}


func (inst*p11d85ef560_monitor_Monitor) getEnableLogDetail(ie application.InjectionExt)bool{
    return ie.GetBool("${dnsss-resolver.monitor.log-detail}")
}



// type p2eec8228d.DnsssResponder in package:github.com/starter-go/dns-server-starter/app/classes/responder
//
// id:com-2eec8228d893edfe-responder-DnsssResponder
// class:class-9ee9f73b8d049a5dcf2cc5fdc03de167-ResolverRegistry
// alias:
// scope:singleton
//
type p2eec8228d8_responder_DnsssResponder struct {
}

func (inst* p2eec8228d8_responder_DnsssResponder) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2eec8228d893edfe-responder-DnsssResponder"
	r.Classes = "class-9ee9f73b8d049a5dcf2cc5fdc03de167-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2eec8228d8_responder_DnsssResponder) new() any {
    return &p2eec8228d.DnsssResponder{}
}

func (inst* p2eec8228d8_responder_DnsssResponder) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2eec8228d.DnsssResponder)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*p2eec8228d8_responder_DnsssResponder) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${dnsss-resolver.responder.enabled}")
}



// type peceb8058a.ForwardResolver in package:github.com/starter-go/dns-server-starter/app/classes/upstream
//
// id:com-eceb8058adebfff3-upstream-ForwardResolver
// class:class-9ee9f73b8d049a5dcf2cc5fdc03de167-ResolverRegistry
// alias:
// scope:singleton
//
type peceb8058ad_upstream_ForwardResolver struct {
}

func (inst* peceb8058ad_upstream_ForwardResolver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-eceb8058adebfff3-upstream-ForwardResolver"
	r.Classes = "class-9ee9f73b8d049a5dcf2cc5fdc03de167-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* peceb8058ad_upstream_ForwardResolver) new() any {
    return &peceb8058a.ForwardResolver{}
}

func (inst* peceb8058ad_upstream_ForwardResolver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*peceb8058a.ForwardResolver)
	nop(ie, com)

	
    com.DnsUpstreamServers = inst.getDnsUpstreamServers(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*peceb8058ad_upstream_ForwardResolver) getDnsUpstreamServers(ie application.InjectionExt)string{
    return ie.GetString("${dnsss-resolver.upstream.servers}")
}


func (inst*peceb8058ad_upstream_ForwardResolver) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${dnsss-resolver.upstream.enabled}")
}


