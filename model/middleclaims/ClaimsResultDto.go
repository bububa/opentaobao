package middleclaims

// ClaimsResultDto 
type ClaimsResultDto struct {
    // 服务工单ID
    ServiceWorkOrderId   int64 `json:"service_work_order_id,omitempty" xml:"service_work_order_id,omitempty"`
    // 主订单号
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 子订单号
    SubOrderId   int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
    // 报案号
    ReportNo   string `json:"report_no,omitempty" xml:"report_no,omitempty"`
    // 收货状态
    TakeGoodsStatus   string `json:"take_goods_status,omitempty" xml:"take_goods_status,omitempty"`
    // 理赔结果
    ClaimsResult   bool `json:"claims_result,omitempty" xml:"claims_result,omitempty"`
    // 理赔拒绝原因
    ClaimsResultDesc   string `json:"claims_result_desc,omitempty" xml:"claims_result_desc,omitempty"`
    // 理赔金额
    ClaimAmount   int64 `json:"claim_amount,omitempty" xml:"claim_amount,omitempty"`
    // 理赔币种
    ClaimCurrency   string `json:"claim_currency,omitempty" xml:"claim_currency,omitempty"`
    // 赔付比例
    CompensationRatio   string `json:"compensation_ratio,omitempty" xml:"compensation_ratio,omitempty"`
    // 包裹状态
    PackageStatus   string `json:"package_status,omitempty" xml:"package_status,omitempty"`
    // 预留扩展Map
    ExtensionMap   *Extensionmap `json:"extension_map,omitempty" xml:"extension_map,omitempty"`
}
