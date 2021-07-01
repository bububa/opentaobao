package flight

// ModifyApproveRequestDto 结构体
type ModifyApproveRequestDto struct {
	// 改签申请单号
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 国内国际标识
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
	// 改签数据
	ChangeList []ModifyItemDto `json:"change_list,omitempty" xml:"change_list>modify_item_dto,omitempty"`
}
