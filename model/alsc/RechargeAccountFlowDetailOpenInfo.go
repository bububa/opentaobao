package alsc

// RechargeAccountFlowDetailOpenInfo 
type RechargeAccountFlowDetailOpenInfo struct {

    // 储值账户id
    AccountId   string `json:"account_id,omitempty"`

    // 操作储值资产信息列表
    AffectedPropertyList   []RechargeAccountPropertyOpenInfo `json:"affected_property_list,omitempty"`

    // 操作后储值资产信息列表
    AfterPropertyList   []RechargeAccountPropertyOpenInfo `json:"after_property_list,omitempty"`

    // 操作前储值资产信息列表
    BeforePropertyList   []RechargeAccountPropertyOpenInfo `json:"before_property_list,omitempty"`

    // 品牌Id
    BrandId   string `json:"brand_id,omitempty"`

    // 交易后剩余总金额
    CurrentValue   int64 `json:"current_value,omitempty"`

    // 储值账户流水id
    FlowId   string `json:"flow_id,omitempty"`

    // 储值相关的交易类型
    FlowType   string `json:"flow_type,omitempty"`

    // 操作人ID
    Operator   string `json:"operator,omitempty"`

    // 操作人
    OperatorName   string `json:"operator_name,omitempty"`

    // 外部订单来源
    OrderSrc   int64 `json:"order_src,omitempty"`

    // 交易时间
    OrderTime   string `json:"order_time,omitempty"`

    // 交易总金额，增加为正数，减少为负数
    OrderValue   int64 `json:"order_value,omitempty"`

    // 外部交易单号id
    OuterOrderId   string `json:"outer_order_id,omitempty"`

    // 理由
    Remark   string `json:"remark,omitempty"`

    // 交易门店ID
    ShopId   string `json:"shop_id,omitempty"`

    // 交易门店名称
    ShopName   string `json:"shop_name,omitempty"`

    // 卡ID
    CardId   string `json:"card_id,omitempty"`

    // 原支付单号
    OuterPayId   string `json:"outer_pay_id,omitempty"`

    // 扩展信息
    ExtInfo   string `json:"ext_info,omitempty"`

}
