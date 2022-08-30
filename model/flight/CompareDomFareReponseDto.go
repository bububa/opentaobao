package flight

// CompareDomFareReponseDto 结构体
type CompareDomFareReponseDto struct {
	// 返回政策信息
	PriceComparisonList []PriceComparisonDto `json:"price_comparison_list,omitempty" xml:"price_comparison_list>price_comparison_dto,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
