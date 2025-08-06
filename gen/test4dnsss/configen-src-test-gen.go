package test4dnsss
import (
    pd9888079a "github.com/starter-go/dns-server-starter/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type pd9888079a.DemoUnit in package:github.com/starter-go/dns-server-starter/src/test/golang/unit
//
// id:com-d9888079a3ec2e2c-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pd9888079a3_unit_DemoUnit struct {
}

func (inst* pd9888079a3_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d9888079a3ec2e2c-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd9888079a3_unit_DemoUnit) new() any {
    return &pd9888079a.DemoUnit{}
}

func (inst* pd9888079a3_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd9888079a.DemoUnit)
	nop(ie, com)

	


    return nil
}


