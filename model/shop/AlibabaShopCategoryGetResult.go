package shop

// AlibabaShopCategoryGetResult 结构体
type AlibabaShopCategoryGetResult struct {
	// 返回错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 分类总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 分类对象
	Module *OpenCategoryDto `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
