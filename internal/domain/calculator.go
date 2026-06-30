package domain

// Frontend — уровень стартового цифрового присутствия (SMM/бренд-база).
type Frontend struct {
	Value       string `json:"value"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

// Backend — стратегический уровень охвата (локальный / европейский / глобальный).
type Backend struct {
	Value            string `json:"value"`
	Label            string `json:"label"`
	Description      string `json:"description"`
	PricePerFunction int    `json:"pricePerFunction"`
}

// Database — тип контент-системы (объём и сложность контент-маркетинга).
type Database struct {
	Value             string   `json:"value"`
	Label             string   `json:"label"`
	Description       string   `json:"description"`
	BasePrice         int      `json:"basePrice"`
	BackendPercentage int      `json:"backendPercentage"`
	Disabled          []string `json:"disabled"`
}

// Infrastructure — архитектура каналов и дистрибуции.
type Infrastructure struct {
	Value      string `json:"value"`
	Label      string `json:"label"`
	Percentage int    `json:"percentage"`
	MinPrice   int    `json:"minPrice"`
	MaxPrice   *int   `json:"maxPrice"`
}

// CalculationInput — входные данные калькулятора EURO SMM.
type CalculationInput struct {
	Frontend       Frontend
	Backend        Backend
	Database       Database
	Infrastructure Infrastructure
	Functions      int // количество ключевых направлений / кампаний
}

// CalculationResult — результат расчёта.
type CalculationResult struct {
	FrontendCost int
	BackendCost  int
	DatabaseCost int
	InfraCost    int
	Total        int
}
