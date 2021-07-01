package icburfq

// Categorylist 结构体
type Categorylist struct {
	// 类目ID
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}
