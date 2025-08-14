package dnsss

import "github.com/miekg/dns"

type Context struct {
	Writer   dns.ResponseWriter
	Request  *dns.Msg
	Response *dns.Msg
}

func (inst *Context) CountAnswer() int {
	if inst == nil {
		return 0
	}
	resp := inst.Response
	if resp == nil {
		return 0
	}
	return len(resp.Answer)
}

func (inst *Context) HasAnswer() bool {
	cnt := inst.CountAnswer()
	return cnt > 0
}
