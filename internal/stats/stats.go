package stats

import (
	"fmt"
	"time"
)

func LogTimeExecution(since time.Time) {
	fmt.Println("Execution time:", time.Since(since))
}
