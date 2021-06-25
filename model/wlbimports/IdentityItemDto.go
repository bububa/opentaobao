package wlbimports

// IdentityItemDto 
type IdentityItemDto struct {

    // 鉴定次数
    IdentityCnt   int64 `json:"identity_cnt,omitempty"`

    // 防伪扣编码
    UniCode   string `json:"uni_code,omitempty"`

    // 货品Id
    ScItemId   string `json:"sc_item_id,omitempty"`

    // 鉴定结果备注
    IdentityRemark   string `json:"identity_remark,omitempty"`

    // 鉴定报告地址
    ReportUrl   string `json:"report_url,omitempty"`

    // 鉴定结果
    IdentityResult   string `json:"identity_result,omitempty"`

}
