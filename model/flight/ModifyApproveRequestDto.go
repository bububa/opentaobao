package flight

// ModifyApproveRequestDto 
/* model for simplify = false
type ModifyApproveRequestDto struct {

    // 改签申请单号
    
    ApplyId   string `json:"apply_id,omitempty"`
    

    // 国内国际标识
    
    DomesticIntl   int64 `json:"domestic_intl,omitempty"`
    

    // 改签数据
    
    ChangeList  struct {
        ModifyItemDTO  []ModifyItemDTO `json:"modify_item_dto,omitempty"`
    } `json:"change_list,omitempty"`
    

}
*/

// ModifyApproveRequestDto 
type ModifyApproveRequestDto struct {

    // 改签申请单号
    ApplyId   string `json:"apply_id,omitempty"`

    // 国内国际标识
    DomesticIntl   int64 `json:"domestic_intl,omitempty"`

    // 改签数据
    ChangeList   []ModifyItemDTO `json:"change_list,omitempty"`

}
