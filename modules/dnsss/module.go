package dnsss

import (
	"github.com/starter-go/application"
	dnsserverstarter "github.com/starter-go/dns-server-starter"
	"github.com/starter-go/dns-server-starter/gen/main4dnsss"
	"github.com/starter-go/dns-server-starter/gen/test4dnsss"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/security/modules/security"
)

// Module  ...
func Module() application.Module {
	mb := dnsserverstarter.NewMainModule()
	mb.Components(main4dnsss.ExportComponents)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(security.Module())

	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := dnsserverstarter.NewTestModule()
	mb.Components(test4dnsss.ExportComponents)
	return mb.Create()
}
