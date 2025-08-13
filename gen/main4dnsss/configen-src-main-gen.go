package main4dnsss
import (
    p9fd301bff "github.com/starter-go/dns-server-starter/app/classes/cache"
    pc47b01c43 "github.com/starter-go/dns-server-starter/app/classes/dnsss"
    p11d85ef56 "github.com/starter-go/dns-server-starter/app/classes/monitor"
    p2eec8228d "github.com/starter-go/dns-server-starter/app/classes/responder"
    peceb8058a "github.com/starter-go/dns-server-starter/app/classes/upstream"
    pd70af8b23 "github.com/starter-go/dns-server-starter/app/data/dao"
    p6365d7957 "github.com/starter-go/dns-server-starter/app/data/database"
    pc9a3f7a59 "github.com/starter-go/dns-server-starter/app/data/dxo"
    p7d9911eb4 "github.com/starter-go/dns-server-starter/app/implements/example"
    p421feda54 "github.com/starter-go/dns-server-starter/app/web/controllers/apiv1"
    p9ee9f73b8 "github.com/starter-go/dns-server-starter/dnsss"
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
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

	


    return nil
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

	


    return nil
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

	


    return nil
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

	


    return nil
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


    return nil
}


func (inst*peceb8058ad_upstream_ForwardResolver) getDnsUpstreamServers(ie application.InjectionExt)string{
    return ie.GetString("${dns.resolver.upstream.servers}")
}



// type p6365d7957.ThisGroup in package:github.com/starter-go/dns-server-starter/app/data/database
//
// id:com-6365d79579708638-database-ThisGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:alias-c9a3f7a595874e44cb593cae1ec88572-DatabaseAgent
// scope:singleton
//
type p6365d79579_database_ThisGroup struct {
}

func (inst* p6365d79579_database_ThisGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6365d79579708638-database-ThisGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = "alias-c9a3f7a595874e44cb593cae1ec88572-DatabaseAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6365d79579_database_ThisGroup) new() any {
    return &p6365d7957.ThisGroup{}
}

func (inst* p6365d79579_database_ThisGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6365d7957.ThisGroup)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Alias = inst.getAlias(ie)
    com.URI = inst.getURI(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)
    com.SourceManager = inst.getSourceManager(ie)


    return nil
}


func (inst*p6365d79579_database_ThisGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.default.enabled}")
}


func (inst*p6365d79579_database_ThisGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.alias}")
}


func (inst*p6365d79579_database_ThisGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.uri}")
}


func (inst*p6365d79579_database_ThisGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.table-name-prefix}")
}


func (inst*p6365d79579_database_ThisGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.datasource}")
}


func (inst*p6365d79579_database_ThisGroup) getSourceManager(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}



// type p7d9911eb4.DaoImpl in package:github.com/starter-go/dns-server-starter/app/implements/example
//
// id:com-7d9911eb42a2a4c3-example-DaoImpl
// class:
// alias:alias-d70af8b23bc2aef64c08d4b7680ff384-ExampleDAO
// scope:singleton
//
type p7d9911eb42_example_DaoImpl struct {
}

func (inst* p7d9911eb42_example_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7d9911eb42a2a4c3-example-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-d70af8b23bc2aef64c08d4b7680ff384-ExampleDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7d9911eb42_example_DaoImpl) new() any {
    return &p7d9911eb4.DaoImpl{}
}

func (inst* p7d9911eb42_example_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7d9911eb4.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)


    return nil
}


func (inst*p7d9911eb42_example_DaoImpl) getAgent(ie application.InjectionExt)pc9a3f7a59.DatabaseAgent{
    return ie.GetComponent("#alias-c9a3f7a595874e44cb593cae1ec88572-DatabaseAgent").(pc9a3f7a59.DatabaseAgent)
}



// type p421feda54.ExampleController in package:github.com/starter-go/dns-server-starter/app/web/controllers/apiv1
//
// id:com-421feda5413d45b5-apiv1-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p421feda541_apiv1_ExampleController struct {
}

func (inst* p421feda541_apiv1_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-421feda5413d45b5-apiv1-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p421feda541_apiv1_ExampleController) new() any {
    return &p421feda54.ExampleController{}
}

func (inst* p421feda541_apiv1_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p421feda54.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p421feda541_apiv1_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p421feda541_apiv1_ExampleController) getDao(ie application.InjectionExt)pd70af8b23.ExampleDAO{
    return ie.GetComponent("#alias-d70af8b23bc2aef64c08d4b7680ff384-ExampleDAO").(pd70af8b23.ExampleDAO)
}


