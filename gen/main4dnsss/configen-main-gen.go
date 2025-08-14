package main4dnsss

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p11d85ef560_monitor_ErrorFilter{})
    inst.register(&p11d85ef560_monitor_Monitor{})
    inst.register(&p2eec8228d8_responder_DnsssResponder{})
    inst.register(&p421feda541_apiv1_ExampleController{})
    inst.register(&p6365d79579_database_ThisGroup{})
    inst.register(&p7d9911eb42_example_DaoImpl{})
    inst.register(&p9fd301bff6_cache_Cache{})
    inst.register(&pc47b01c433_dnsss_ServerStarter{})
    inst.register(&peceb8058ad_upstream_ForwardResolver{})


    return nil
}
