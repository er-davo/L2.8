package myntp

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// PrintCurrentTime prints the current time
func PrintCurrentTime() {
	curTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(curTime.Format(time.Layout))
}
