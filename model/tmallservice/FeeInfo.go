package tmallservice

// FeeInfo 
type FeeInfo struct {

    // 金额单价(分)
    Amount   string `json:"amount,omitempty"`

    // 出资方code
    FromRoleCode   string `json:"from_role_code,omitempty"`

    // 收款方的id
    ToRoleId   int64 `json:"to_role_id,omitempty"`

    // 费用项科目code
    ItemCode   string `json:"item_code,omitempty"`

    // 出资方id
    FromRoleId   int64 `json:"from_role_id,omitempty"`

    // 收款方的code
    ToRoleCode   string `json:"to_role_code,omitempty"`

}
