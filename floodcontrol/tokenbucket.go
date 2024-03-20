package floodcontrol

import (
	"fmt"
	"math"
	"time"
)

func (tb *TokenBucket) RefillBucket() {
	now := time.Now()
	tim := time.Duration(tb.timeToRefill) * time.Second
	prev := now.Add(-tim)
	if prev.Before(tb.lastRefillTime) {
		fmt.Println("before")
		tb.tokens = int(math.Min(float64(tb.tokens+tb.refillTokens), float64(tb.maxTokens)))
		tb.lastRefillTime = now
	}
}

func (tb *TokenBucket) Request() bool {
	tb.RefillBucket()
	if tb.tokens > 0 {
		tb.tokens -= 1
		return true
	}
	return false
}
