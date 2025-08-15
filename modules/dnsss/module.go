package dnsss

import (
	"github.com/starter-go/application"
	dnsserverstarter "github.com/starter-go/dns-server-starter"
	"github.com/starter-go/dns-server-starter/gen/main4dnsss"
	"github.com/starter-go/dns-server-starter/gen/test4dnsss"
	"github.com/starter-go/starter"
	"github.com/starter-go/stopper/modules/stopper"
)

// Module  ...
func Module() application.Module {
	mb := dnsserverstarter.NewMainModule()
	mb.Components(main4dnsss.ExportComponents)

	// mb.Depend(libgin.Module())
	// mb.Depend(security.Module())
	// mb.Depend(libgorm.Module())
	// mb.Depend(mysql.Module())
	// mb.Depend(sqlserver.Module())

	mb.Depend(starter.Module())
	mb.Depend(stopper.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := dnsserverstarter.NewTestModule()
	mb.Components(test4dnsss.ExportComponents)
	return mb.Create()
}
