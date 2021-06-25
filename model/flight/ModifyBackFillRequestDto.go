package flight

// ModifyBackFillRequestDto 
type ModifyBackFillRequestDto struct {

    // 申请单号
    ApplyId   string `json:"apply_id,omitempty"`

    // 国际国内标识
    DomesticIntl   int64 `json:"domestic_intl,omitempty"`

    // 改签数据
    ChangeList   []ModifyItemDTO `json:"change_list,omitempty"`

    // 币种
    Currency   string `json:"currency,omitempty"`

}
