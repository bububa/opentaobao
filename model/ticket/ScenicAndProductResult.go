package ticket

// ScenicAndProductResult 结构体
type ScenicAndProductResult struct {
	// 景点列表
	ScenicList []Scenic `json:"scenic_list,omitempty" xml:"scenic_list>scenic,omitempty"`
}
