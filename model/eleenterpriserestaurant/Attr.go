package eleenterpriserestaurant

// Attr 结构体
type Attr struct {
	// 值
	Values []string `json:"values,omitempty" xml:"values>string,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
