package eleenterpriserestaurant

// Specification 结构体
type Specification struct {
	// 特别说明
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 规格说明
	Values []string `json:"values,omitempty" xml:"values>string,omitempty"`
}
