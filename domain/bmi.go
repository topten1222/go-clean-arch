package domain

type BmiRequest struct {
	Weight float64 `json:"weight" validate:"required"`
	Height float64 `json:"height" validate:"required"`
}

type BmiReponse struct {
	Score float64 `json:"score"`
	Name  string  `json:"name"`
}
