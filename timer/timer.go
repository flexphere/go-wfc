package timer

import (
	"fmt"
	"time"
)

var timers = make(map[string]time.Time)

func Start(name string) {
	timers[name] = time.Now()
}

func End(name string) {
	if val, ok := timers[name]; ok {
		nano := time.Since(val).Nanoseconds()
		fmt.Printf("%s: %.2fms\n", name, float64(nano)/1000000)
	}
}
