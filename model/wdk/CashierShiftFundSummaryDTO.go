package wdk

// CashierShiftFundSummaryDTO 
type CashierShiftFundSummaryDTO struct {
    // 商户编号
    MerchantCode   string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
    // 门店编号
    ShopCode   string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
    // 业务日期
    BizDate   string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
    // POS机编号
    PosNo   string `json:"pos_no,omitempty" xml:"pos_no,omitempty"`
    // 值班编号
    ShiftNo   string `json:"shift_no,omitempty" xml:"shift_no,omitempty"`
    // 收银员姓名
    CashierName   string `json:"cashier_name,omitempty" xml:"cashier_name,omitempty"`
    // 支付方式
    PayMethod   string `json:"pay_method,omitempty" xml:"pay_method,omitempty"`
    // 系统统计金额
    SystemAmount   string `json:"system_amount,omitempty" xml:"system_amount,omitempty"`
    // 输入金额
    InputAmount   string `json:"input_amount,omitempty" xml:"input_amount,omitempty"`
    // 单据状态
    SettleStatus   string `json:"settle_status,omitempty" xml:"settle_status,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 扩展字段
    Extra   string `json:"extra,omitempty" xml:"extra,omitempty"`
}
