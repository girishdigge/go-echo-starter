package repository

import "github.com/girishdigge/go-echo-starter/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
