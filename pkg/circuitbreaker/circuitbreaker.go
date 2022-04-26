package circuitbreaker

import (
	"fmt"
	"time"

	test "bitbucket.org/panda-development/go-package"
	"github.com/sony/gobreaker"
)

var CB *gobreaker.CircuitBreaker

func Init() {
	test.Print()
	cfg := gobreaker.Settings{
		// When to flush counters int the Closed state
		Interval: 10 * time.Second,
		// Time to switch from Open to Half-open
		Timeout: 10 * time.Second,
		// Function with check when to switch from Closed to Open
		ReadyToTrip: shouldBeSwitchedToOpen,
		OnStateChange: func(_ string, from gobreaker.State, to gobreaker.State) {
			fmt.Println("state changed from", from.String(), "to", to.String())
		},
	}

	CB = gobreaker.NewCircuitBreaker(cfg)
}

func shouldBeSwitchedToOpen(counts gobreaker.Counts) bool {
	failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
	return counts.Requests >= 5 && failureRatio >= 0.6
}
