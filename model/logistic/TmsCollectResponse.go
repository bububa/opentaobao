package logistic

// TmsCollectResponse 结构体
type TmsCollectResponse struct {
	// 风控通过重量
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 重量是否通过风控
	WeightCheckResult bool `json:"weight_check_result,omitempty" xml:"weight_check_result,omitempty"`
}
