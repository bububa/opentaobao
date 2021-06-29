package koubeimall

// MallDetailDTO 
type MallDetailDTO struct {
    // 商圈内门店类目模型
    CategoryTabInfoList   []CategoryTabInfoDTO `json:"category_tab_info_list,omitempty" xml:"category_tab_info_list>category_tab_info_dto,omitempty"`
    // 商圈信息模型
    MallDto   *MallDTO `json:"mall_dto,omitempty" xml:"mall_dto,omitempty"`
}
