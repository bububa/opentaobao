package iot

// PageInfo 结构体
type PageInfo struct {
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 分页页码
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 数据集
	List []BusinessRecipeOpenDto `json:"list,omitempty" xml:"list>business_recipe_open_dto,omitempty"`
}
