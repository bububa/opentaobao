package flight

// ModifyBackFillRequestDto 
/* model for simplify = false
type ModifyBackFillRequestDto struct {

    // 申请单号
    
    ApplyId   string `json:"apply_id,omitempty"`
    

    // 国际国内标识
    
    DomesticIntl   int64 `json:"domestic_intl,omitempty"`
    

    // 改签数据
    
    ChangeList  struct {
        ModifyItemDTO  []ModifyItemDTO `json:"modify_item_dto,omitempty"`
    } `json:"change_list,omitempty"`
    

    // 币种
    
    Currency   string `json:"currency,omitempty"`
    

}
*/

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
