package koubeimall

// CategoryTabInfoDto 结构体
type CategoryTabInfoDto struct {
	// 前台类目ids
	CategoryIdList []string `json:"category_id_list,omitempty" xml:"category_id_list>string,omitempty"`
	// 前台类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
}
