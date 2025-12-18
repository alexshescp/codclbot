package calculator

import (
	"codclbot/internal/config"
	"codclbot/internal/domain"
)

type Service struct {
	Pricing *config.Pricing
}

func NewService(p *config.Pricing) *Service {
	return &Service{Pricing: p}
}

func (s *Service) CalculateForInput(in domain.CalculationInput) domain.CalculationResult {
	return domain.CalculateCost(in)
}
