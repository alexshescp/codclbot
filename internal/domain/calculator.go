package domain

// Frontend — тип фронтенд-входа (один в один с JSON-прайсом)
type Frontend struct {
	Value       string `json:"value"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Price       int    `json:"price"` // EUR
}

// Backend — тип backend-варианта
type Backend struct {
	Value            string `json:"value"`
	Label            string `json:"label"`
	Description      string `json:"description"`
	PricePerFunction int    `json:"pricePerFunction"` // EUR / функция
}

// Database — тип варианта БД
type Database struct {
	Value             string   `json:"value"`
	Label             string   `json:"label"`
	Description       string   `json:"description"`
	BasePrice         int      `json:"basePrice"`         // EUR
	BackendPercentage int      `json:"backendPercentage"` // % от backendCost
	Disabled          []string `json:"disabled"`          // backend.value, для которых нельзя
}

// Infrastructure — инфраструктура (облако/виртуалка/шаред/без инфры)
type Infrastructure struct {
	Value      string `json:"value"`
	Label      string `json:"label"`
	Percentage int    `json:"percentage"` // % от (frontend+backend+db)
	MinPrice   int    `json:"minPrice"`   // минимальная стоимость, EUR
	MaxPrice   *int   `json:"maxPrice"`   // максимум, если задан
}

// CalculationInput — всё, что нужно, чтобы посчитать итог
type CalculationInput struct {
	Frontend       Frontend
	Backend        Backend
	Database       Database
	Infrastructure Infrastructure
	Functions      int
}

// CalculationResult — результат расчёта
type CalculationResult struct {
	FrontendCost int
	BackendCost  int
	DatabaseCost int
	InfraCost    int
	Total        int
}

// CalculateCost — чистый доменный расчёт стоимости.
// Никакого Telegram внутри, только математика.
func CalculateCost(in CalculationInput) CalculationResult {
	frontendCost := in.Frontend.Price
	backendCost := in.Backend.PricePerFunction * in.Functions
	dbCost := in.Database.BasePrice + (backendCost*in.Database.BackendPercentage)/100

	base := frontendCost + backendCost + dbCost

	infraCost := 0
	if in.Infrastructure.Percentage > 0 {
		infraCost = (base * in.Infrastructure.Percentage) / 100
		if infraCost < in.Infrastructure.MinPrice {
			infraCost = in.Infrastructure.MinPrice
		}
		if in.Infrastructure.MaxPrice != nil && infraCost > *in.Infrastructure.MaxPrice {
			infraCost = *in.Infrastructure.MaxPrice
		}
	}

	total := frontendCost + backendCost + dbCost + infraCost

	return CalculationResult{
		FrontendCost: frontendCost,
		BackendCost:  backendCost,
		DatabaseCost: dbCost,
		InfraCost:    infraCost,
		Total:        total,
	}
}
