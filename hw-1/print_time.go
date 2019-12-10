package print_time

import (
	"fmt"

	"github.com/beevik/ntp"
)

func Print_dt () {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if (if ) {
		fmt.Printf("time now is %s\n", time)
	} else {
		return false
	}
	return true
}