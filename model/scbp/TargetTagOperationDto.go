package scbp

// TargetTagOperationDto 结构体
type TargetTagOperationDto struct {
	// crowd or region
	Scope []string `json:"scope,omitempty" xml:"scope>string,omitempty"`
	// 出价类型：0=出价, 1=溢价，2=过滤, 3=召回
	PriceMode int64 `json:"price_mode,omitempty" xml:"price_mode,omitempty"`
}
