package eleenterprisecartnew

// EnterpriseData 
type EnterpriseData struct {
    // 购物车原始金额
    OriginalTotal   string `json:"original_total,omitempty" xml:"original_total,omitempty"`
    // 购物车金额
    Total   string `json:"total,omitempty" xml:"total,omitempty"`
    // 创建购物车时间戳
    CreateTime   int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 手机号
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    // 额外信息
    Extras   []Extra `json:"extras,omitempty" xml:"extras>extra,omitempty"`
    // 费用说明
    ServiceFeeExplanation   string `json:"service_fee_explanation,omitempty" xml:"service_fee_explanation,omitempty"`
    // 起送价
    DeliverAmount   string `json:"deliver_amount,omitempty" xml:"deliver_amount,omitempty"`
    // 购物车篮子
    Groups   []string `json:"groups,omitempty" xml:"groups>string,omitempty"`
    // 购物车id
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
}
