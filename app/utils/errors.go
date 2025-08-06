package utils

import "github.com/starter-go/vlog"

func HandleError(err error) {
	if err == nil {
		return
	}
	vlog.Error("%s", err.Error())
}

func HandleErrorX(x any) {
	if x == nil {
		return
	}
	err, ok := x.(error)
	if ok {
		HandleError(err)
		return
	}
	vlog.Error("unknown panic: %s", x)
}
