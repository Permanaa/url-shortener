package limiter

import (
	"errors"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity   float64   // Maximum number of tokens the bucket can hold
	rate       float64   // Number of tokens to add per second
	tokens     float64   // Current number of tokens in the bucket
	lastRefill time.Time // Timestamp of the last token refill
	mutex      sync.Mutex
}

func NewTokenBucket(capacity, rate float64) (*TokenBucket, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity must be greater than zero")
	}
	if rate <= 0 {
		return nil, errors.New("rate must be greater than zero")
	}
	return &TokenBucket{
		capacity:   float64(capacity),
		rate:       float64(rate),
		tokens:     float64(capacity),
		lastRefill: time.Now(),
	}, nil
}

func (tb *TokenBucket) Take(tokens int) bool {
	if tokens <= 0 {
		return false
	}

	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	tb.refill()

	requested := float64(tokens)
	if tb.tokens >= requested {
		tb.tokens -= requested
		return true
	}

	return false
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()

	tb.tokens += elapsed * tb.rate

	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	tb.lastRefill = now
}
