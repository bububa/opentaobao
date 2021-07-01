package drugtrace

// BillUpstreamDto 
type BillUpstreamDto struct {
    // 发货企业名称
    FromUserName   string `json:"from_user_name,omitempty" xml:"from_user_name,omitempty"`
    // 单据时间
    BillTime   string `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
    // 货主
    RefUserId   string `json:"ref_user_id,omitempty" xml:"ref_user_id,omitempty"`
    // 发货企业ID
    FromUserId   string `json:"from_user_id,omitempty" xml:"from_user_id,omitempty"`
    // 单据类型
    BillType   string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
    // 收货企业名称
    ToUserName   string `json:"to_user_name,omitempty" xml:"to_user_name,omitempty"`
    // 单据号
    BillCode   string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
    // 收货企业ID
    ToUserId   string `json:"to_user_id,omitempty" xml:"to_user_id,omitempty"`
    // 收货企业REF_ENT_ID
    ToRefUserId   string `json:"to_ref_user_id,omitempty" xml:"to_ref_user_id,omitempty"`
    // 发货企业REF_ENT_ID
    FromRefUserId   string `json:"from_ref_user_id,omitempty" xml:"from_ref_user_id,omitempty"`
}
