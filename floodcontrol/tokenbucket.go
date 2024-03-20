package floodcontrol

import (
	"math"
	"os"
	"strconv"
	"time"
)

// NewTokenBucket констуктор для TokenBucket
func NewTokenBucket(tokens int, lastRefillTime time.Time) *TokenBucket {
	maxTokens, _ := strconv.Atoi(os.Getenv("TOKENS"))
	refillTokens := maxTokens
	timeToRefill, _ := strconv.Atoi(os.Getenv("SECONDS"))
	return &TokenBucket{
		tokens:         tokens,
		maxTokens:      maxTokens,
		refillTokens:   refillTokens,
		lastRefillTime: lastRefillTime,
		timeToRefill:   timeToRefill,
	}
}

// refillBucket пополняет ведро токенами если прошло время указанное в timeToRefill
func (tb *TokenBucket) refillBucket() {
	now := time.Now()
	pass := time.Duration(tb.timeToRefill) * time.Second
	prevTime := now.Add(-pass)
	if !prevTime.Before(tb.lastRefillTime) || prevTime.Equal(tb.lastRefillTime) {
		tb.tokens = int(math.Min(float64(tb.tokens+tb.refillTokens), float64(tb.maxTokens)))
		tb.lastRefillTime = now
	}
}

// request в случае если токены еще остались отнимает их количество иначе - пользователь флудит
func (tb *TokenBucket) request() bool {
	tb.refillBucket()
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}
