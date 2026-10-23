package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

type Cors struct {
	cors *cors.Cors
}

func NewCors() *Cors {
	return &Cors{
		cors: cors.New(cors.Options{
			AllowedOrigins: []string{
				"http://localhost:3000",
			},
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
				http.MethodOptions,
			},
			AllowedHeaders: []string{
				"Content-Type",
				"Authorization",
			},
			AllowCredentials: true,
		}),
	}
}

func (c *Cors) Apply(next http.Handler) http.Handler {
	return c.cors.Handler(next)
}
