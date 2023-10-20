package shop

// AlibabaShopCategoryAllGetResult 结构体
type AlibabaShopCategoryAllGetResult struct {
	// 分类对象
	ModuleList []OpenCategoryDto `json:"module_list,omitempty" xml:"module_list>open_category_dto,omitempty"`
	// 返回错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 分类总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
