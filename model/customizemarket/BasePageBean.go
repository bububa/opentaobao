package customizemarket

// BasePageBean 结构体
type BasePageBean struct {
	// 数据
	Data []PetProfileDto `json:"data,omitempty" xml:"data>pet_profile_dto,omitempty"`
	// 当前页面
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总行数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
