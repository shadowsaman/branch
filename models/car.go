package models

type CarPrimeryKey struct {
	Id string `json:"id"`
}

type CreateCar struct {
	StateNumber        string  `json:"state_number"`
	Model              string  `json:"model"`
	Price              float64 `json:"price"`
	DailyLimit         int     `json:"daily_limit"`
	OverLimit          int     `json:"over_limit"`
	InvestorPercentage float64 `json:"investor_percentage"`
	InvestorId         string  `json:"investor_id"`
	BranchId           string  `json:"branch_id"`
	Km                 int     `json:"km"`
}

type Car struct {
	Id                 string  `json:"id"`
	StateNumber        string  `json:"state_number"`
	Model              string  `json:"model"`
	Status             string  `json:"status"`
	Price              float64 `json:"price"`
	DailyLimit         int     `json:"daily_limit"`
	OverLimit          int     `json:"over_limit"`
	InvestorPercentage float64 `json:"investor_percentage"`
	InvestorId         string  `json:"investor_id"`
	BranchId           string  `json:"branch_id"`
	Km                 int     `json:"km"`
	CreatedAt          string  `json:"crated_at"`
	UpdatedAt          string  `json:"updated_at"`
}

type UpdateCar struct {
	Id                 string  `json:"id"`
	StateNumber        string  `json:"state_number"`
	Model              string  `json:"model"`
	Status             string  `json:"status"`
	Price              float64 `json:"price"`
	DailyLimit         int     `json:"daily_limit"`
	OverLimit          int     `json:"over_limit"`
	InvestorPercentage float64 `json:"investor_percentage"`
	InvestorId         string  `json:"investor_id"`
	BranchId           string  `json:"branch_id"`
	Km                 int     `json:"km"`
}

type UpdateCarSwag struct {
	StateNumber        string  `json:"state_number"`
	Model              string  `json:"model"`
	Status             string  `json:"status"`
	Price              float64 `json:"price"`
	DailyLimit         int     `json:"daily_limit"`
	OverLimit          int     `json:"over_limit"`
	InvestorPercentage float64 `json:"investor_percentage"`
	InvestorId         string  `json:"investor_id"`
	Km                 int     `json:"km"`
	BranchId           string  `json:"branch_id"`
}

type GetListCarRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListCarResponse struct {
	Count int64  `json:"count"`
	Cars  []*Car `json:"cars"`
}
