package flight

// ModifyApproveRequestDto 结构体
type ModifyApproveRequestDto struct {
	// 改签数据
	ChangeList []ModifyItemDto `json:"change_list,omitempty" xml:"change_list>modify_item_dto,omitempty"`
	// 改签申请单号
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 国内国际标识:1:国内,2:国际
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
	// 0:原路退回; 1:退银行卡; 2:原路退回+退银行卡
	RefundWayType int64 `json:"refund_way_type,omitempty" xml:"refund_way_type,omitempty"`
}
