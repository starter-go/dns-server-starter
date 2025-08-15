package test4dnsss
import (
    p3403d6cb1 "github.com/starter-go/dns-server-starter/src/test/golang/units4dnsss"
     "github.com/starter-go/application"
)

// type p3403d6cb1.DemoUnit in package:github.com/starter-go/dns-server-starter/src/test/golang/units4dnsss
//
// id:com-3403d6cb189096c1-units4dnsss-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p3403d6cb18_units4dnsss_DemoUnit struct {
}

func (inst* p3403d6cb18_units4dnsss_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3403d6cb189096c1-units4dnsss-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3403d6cb18_units4dnsss_DemoUnit) new() any {
    return &p3403d6cb1.DemoUnit{}
}

func (inst* p3403d6cb18_units4dnsss_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3403d6cb1.DemoUnit)
	nop(ie, com)

	


    return nil
}


