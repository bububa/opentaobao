package flight

// ModifyBackFillRequestDto 结构体
type ModifyBackFillRequestDto struct {
	// 改签数据
	ChangeList []ModifyItemDto `json:"change_list,omitempty" xml:"change_list>modify_item_dto,omitempty"`
	// 申请单号
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 国际国内标识:1:国内,2:国际
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
}
