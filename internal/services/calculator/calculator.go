package calculator

import (
	"codclbot/internal/domain"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) CalculateForInput(in domain.CalculationInput) domain.CalculationResult {
	// Визуальная и брендовая база
	frontendCost := in.Frontend.Price

	// Стратегический уровень охвата, масштабируется количеством кампаний
	backendCost := in.Backend.PricePerFunction * in.Functions

	// Контент-система: базовая цена + процент от стратегического слоя
	dbCost := in.Database.BasePrice
	dbCost += backendCost * in.Database.BackendPercentage / 100

	// Архитектура каналов и дистрибуции
	infraCost := 0
	if in.Infrastructure.Percentage > 0 {
		infraCost = (frontendCost + backendCost + dbCost) * in.Infrastructure.Percentage / 100
		if infraCost < in.Infrastructure.MinPrice {
			infraCost = in.Infrastructure.MinPrice
		}
		if in.Infrastructure.MaxPrice != nil && infraCost > *in.Infrastructure.MaxPrice {
			infraCost = *in.Infrastructure.MaxPrice
		}
	}

	total := frontendCost + backendCost + dbCost + infraCost

	return domain.CalculationResult{
		FrontendCost: frontendCost,
		BackendCost:  backendCost,
		DatabaseCost: dbCost,
		InfraCost:    infraCost,
		Total:        total,
	}
}
