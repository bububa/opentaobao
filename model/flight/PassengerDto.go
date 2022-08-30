package flight

// PassengerDto 结构体
type PassengerDto struct {
	// 证件类型
	DocumentsType []string `json:"documents_type,omitempty" xml:"documents_type>string,omitempty"`
	// 年龄限制
	AgeLimit string `json:"age_limit,omitempty" xml:"age_limit,omitempty"`
	// 身份地域限制
	DocumentsLimit string `json:"documents_limit,omitempty" xml:"documents_limit,omitempty"`
	// 任务限制
	PaxNum string `json:"pax_num,omitempty" xml:"pax_num,omitempty"`
	// 产品类型
	ProductCode int64 `json:"product_code,omitempty" xml:"product_code,omitempty"`
}
