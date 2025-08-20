package dnsss

import "github.com/starter-go/vlog"

// LoggerAgent: 这个是 dnsss 专用的日志代理接口, 它代理了 dnsss 中各个组件的日志工作, 并委托给 vlog 执行
type LoggerAgent interface {
	GetLogger() vlog.Logger
}
