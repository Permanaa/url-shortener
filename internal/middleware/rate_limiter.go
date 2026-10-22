package middleware

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/permanaa/url-shortener/internal/helper"
	"github.com/permanaa/url-shortener/internal/limiter"
)

type Visitor struct {
	limiter  *limiter.TokenBucket
	lastSeen time.Time
}

type RateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*Visitor
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		visitors: make(map[string]*Visitor),
	}
}

func (i *RateLimiter) Apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getClientIp(r)
		limiter := i.GetLimiter(ip)

		if !limiter.Take(1) {
			helper.JsonError(w, http.StatusTooManyRequests, "rate limit exceeded")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (i *RateLimiter) GetLimiter(ip string) *limiter.TokenBucket {
	i.mu.Lock()
	defer i.mu.Unlock()

	client, exists := i.visitors[ip]
	if !exists {
		limiter, err := limiter.NewTokenBucket(5, 1.0/6.0) // capacity 5, replenish 1 token per 6 seconds
		if err != nil {
			log.Println("err new limiter:", ip, err)
		}
		client = &Visitor{
			limiter: limiter,
		}

		i.visitors[ip] = client
	}

	client.lastSeen = time.Now()

	return client.limiter
}

func getClientIp(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
