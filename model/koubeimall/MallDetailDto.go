package koubeimall

// MallDetailDto 结构体
type MallDetailDto struct {
	// 商圈内门店类目模型
	CategoryTabInfoList []CategoryTabInfoDto `json:"category_tab_info_list,omitempty" xml:"category_tab_info_list>category_tab_info_dto,omitempty"`
	// 商圈信息模型
	MallDto *MallDto `json:"mall_dto,omitempty" xml:"mall_dto,omitempty"`
}
