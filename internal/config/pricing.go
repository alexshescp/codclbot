package config

import (
	"codclbot/internal/bot/session"
	"codclbot/internal/domain"
)

type Pricing struct {
	Frontend       []domain.Frontend
	Backend        []domain.Backend
	Database       []domain.Database
	Infrastructure []domain.Infrastructure
}

func LoadPricing() (*Pricing, error) {
	p := &Pricing{}

	if err := loadFile("internal/config/frontend.json", &p.Frontend); err != nil {
		return nil, err
	}
	if err := loadFile("internal/config/backend.json", &p.Backend); err != nil {
		return nil, err
	}
	if err := loadFile("internal/config/database.json", &p.Database); err != nil {
		return nil, err
	}
	if err := loadFile("internal/config/infrastructure.json", &p.Infrastructure); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Pricing) FindFrontend(label string) *domain.Frontend {
	for _, f := range p.Frontend {
		if f.Label == label {
			return &f
		}
	}
	return nil
}

func (p *Pricing) FindBackend(label string) *domain.Backend {
	for _, f := range p.Backend {
		if f.Label == label {
			return &f
		}
	}
	return nil
}

func (p *Pricing) FindDatabase(label string) *domain.Database {
	for _, f := range p.Database {
		if f.Label == label {
			return &f
		}
	}
	return nil
}

func (p *Pricing) FindInfra(label string) *domain.Infrastructure {
	for _, f := range p.Infrastructure {
		if f.Label == label {
			return &f
		}
	}
	return nil
}

func (p *Pricing) ToCalcInput(sess *session.Session) domain.CalculationInput {
	return domain.CalculationInput{
		Frontend:       *p.FindFrontend(sess.Frontend),
		Backend:        *p.FindBackend(sess.Backend),
		Database:       *p.FindDatabase(sess.Database),
		Infrastructure: *p.FindInfra(sess.Infrastructure),
		Functions:      sess.Functions,
	}
}
