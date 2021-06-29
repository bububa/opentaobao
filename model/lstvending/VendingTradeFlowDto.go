package lstvending

// VendingTradeFlowDto 
type VendingTradeFlowDto struct {
    // 修改时间
    GmtModified   int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 交易类型：1购买，2退款
    TradeType   int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
    // 创建时间
    GmtCreate   int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 设备厂商编码
    SupplierCode   string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
    // 折扣
    Discount   int64 `json:"discount,omitempty" xml:"discount,omitempty"`
    // 外部系统交易流水号
    TradeFlowNo   string `json:"trade_flow_no,omitempty" xml:"trade_flow_no,omitempty"`
    // 外部系统正向交易流水号
    PaymentTradeFlowNo   string `json:"payment_trade_flow_no,omitempty" xml:"payment_trade_flow_no,omitempty"`
    // 设备编码
    EquipmentCode   string `json:"equipment_code,omitempty" xml:"equipment_code,omitempty"`
    // 商品清单
    GoodsDetailDTOList   []VendingTradeGoodsDetailDto `json:"goods_detail_d_t_o_list,omitempty" xml:"goods_detail_d_t_o_list>vending_trade_goods_detail_dto,omitempty"`
    // 交易总金额
    TotalAmount   int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    // 设备ID
    EquipmentId   int64 `json:"equipment_id,omitempty" xml:"equipment_id,omitempty"`
    // 佣金金额
    Commission   int64 `json:"commission,omitempty" xml:"commission,omitempty"`
    // 实际总金额
    ActualAmount   int64 `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
    // 支付明细
    PaymentDTOList   []VendingTradePaymentDto `json:"payment_d_t_o_list,omitempty" xml:"payment_d_t_o_list>vending_trade_payment_dto,omitempty"`
    // 扩展信息
    ExtFields   string `json:"ext_fields,omitempty" xml:"ext_fields,omitempty"`
}
