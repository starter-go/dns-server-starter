package iloggers

import (
	"github.com/starter-go/dns-server-starter/dnsss"
	"github.com/starter-go/vlog"
)

type DnsssLoggerImpl struct {

	//starter:component

	_as func(dnsss.LoggerAgent) //starter:as("#")

	LogDetails bool //starter:inject("${dnsss.logger.log-details}")

	adapter vlog.LoggerAdapter
	logger  vlog.Logger
}

func (inst *DnsssLoggerImpl) _impl() dnsss.LoggerAgent {
	return inst
}

func (inst *DnsssLoggerImpl) GetLogger() vlog.Logger {
	l := inst.logger
	if l == nil {
		l = inst.innerInitLogger()
		inst.logger = l
	}
	return l
}

func (inst *DnsssLoggerImpl) innerGetLevel() vlog.Level {
	if inst.LogDetails {
		return vlog.TRACE
	}
	return vlog.INFO
}

func (inst *DnsssLoggerImpl) innerInitLogger() vlog.Logger {

	lv := inst.innerGetLevel()
	l := &inst.adapter

	l.SetSender(inst)
	l.SetTag("dnsss-logger")
	l.SetLevelAccepted(lv)
	l.SetTargetHandler(vlog.GetLogger())

	return l
}
