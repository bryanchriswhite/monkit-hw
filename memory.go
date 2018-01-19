package hw

import (
	gosigar "github.com/cloudfoundry/gosigar"
	"gopkg.in/spacemonkeygo/monkit.v2"
)

func Memory() monkit.StatSource {
	return monkit.StatSourceFunc(func(cb func(name string, val float64)) {
		var mem gosigar.Mem
		err := mem.Get()
		if err != nil {
			logger.Debuge(err)
			return
		}
		monkit.StatSourceFromStruct(&mem).Stats(cb)
	})
}

func init() {
	registrations["memory"] = Memory()
}