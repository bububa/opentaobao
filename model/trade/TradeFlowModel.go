package trade

// TradeFlowModel 
type TradeFlowModel struct {

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 厂商设备唯一编码
    
    EquipmentCode   string `json:"equipment_code,omitempty" xml:"equipment_code,omitempty"`
    

    // 外部系统正向交易流水号
    
    PaymentTradeFlowNo   string `json:"payment_trade_flow_no,omitempty" xml:"payment_trade_flow_no,omitempty"`
    

    // 实付金额(单位:分)
    
    ActualAmount   int64 `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
    

    // 折扣
    
    Discount   int64 `json:"discount,omitempty" xml:"discount,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 交易流水信息-商品详情
    
    TradeFlowGoodsDetailModelList   []TradeFlowGoodsDetailModel `json:"trade_flow_goods_detail_model_list,omitempty" xml:"trade_flow_goods_detail_model_list,omitempty"`
    

    // 设备类型
    
    EquipmentType   string `json:"equipment_type,omitempty" xml:"equipment_type,omitempty"`
    

    // 支付信息
    
    TradeFlowPaymentModelList   []TradeFlowPaymentModel `json:"trade_flow_payment_model_list,omitempty" xml:"trade_flow_payment_model_list,omitempty"`
    

    // 经营者userId
    
    OperatorUserId   string `json:"operator_user_id,omitempty" xml:"operator_user_id,omitempty"`
    

    // 交易总额(单位:分)
    
    TotalAmount   int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    

    // 经营者名字userId
    
    OperatorUserName   string `json:"operator_user_name,omitempty" xml:"operator_user_name,omitempty"`
    

    // 外部系统交易流水号
    
    TradeFlowNo   string `json:"trade_flow_no,omitempty" xml:"trade_flow_no,omitempty"`
    

    // 设备名称
    
    EquipmentName   string `json:"equipment_name,omitempty" xml:"equipment_name,omitempty"`
    

    // 佣金金额(单位:分)
    
    Commission   int64 `json:"commission,omitempty" xml:"commission,omitempty"`
    

    // 交易类型：1购买，2退款
    
    TradeType   int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
    

}
