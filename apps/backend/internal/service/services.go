package service

import (
	"github.com/girishdigge/go-echo-starter/internal/lib/job"
	"github.com/girishdigge/go-echo-starter/internal/repository"
	"github.com/girishdigge/go-echo-starter/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
