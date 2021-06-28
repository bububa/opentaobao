package waybill

// WaybillApplySubscriptionInfo 
/* model for simplify = false
type WaybillApplySubscriptionInfo struct {

    // CP网点信息及对应的商家的发货信息
    
    BranchAccountCols  struct {
        WaybillBranchAccount  []WaybillBranchAccount `json:"waybill_branch_account,omitempty"`
    } `json:"branch_account_cols,omitempty"`
    

    // 物流服务商ID
    
    CpCode   string `json:"cp_code,omitempty"`
    

    // 物流服务商业务类型 1：直营 2：加盟 3：落地配 4：直营带网点
    
    CpType   int64 `json:"cp_type,omitempty"`
    

}
*/

// WaybillApplySubscriptionInfo 
type WaybillApplySubscriptionInfo struct {

    // CP网点信息及对应的商家的发货信息
    BranchAccountCols   []WaybillBranchAccount `json:"branch_account_cols,omitempty"`

    // 物流服务商ID
    CpCode   string `json:"cp_code,omitempty"`

    // 物流服务商业务类型 1：直营 2：加盟 3：落地配 4：直营带网点
    CpType   int64 `json:"cp_type,omitempty"`

}
