package moscm

// RmaCriteria 
type RmaCriteria struct {
    // 订单号
    OrderNumber   string `json:"order_number,omitempty" xml:"order_number,omitempty"`
    // 退换货单据号
    RmaNumbers   []string `json:"rma_numbers,omitempty" xml:"rma_numbers>string,omitempty"`
    // 单据创建时间范围：结束时间
    EndDate   string `json:"end_date,omitempty" xml:"end_date,omitempty"`
    // 单据创建时间范围：开始时间
    StartDate   string `json:"start_date,omitempty" xml:"start_date,omitempty"`
    // 单据类型，可选值：退货("GOODRETURN"),换货("GOODEXCHANGE"),仅退款（"RETURN"）
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    // 单据状态，可选值：已创建("CREATED"),已收货同意退款("INBOUND"),已收货不同意退款("NOTINBOUND"),已退款("REFUNDED"),已完成("COMPLETED),已作废("Obsolete")
    Status   []string `json:"status,omitempty" xml:"status>string,omitempty"`
    // 供应商专柜Id
    OutCounterId   string `json:"out_counter_id,omitempty" xml:"out_counter_id,omitempty"`
    // 银泰专柜Id
    CounterId   string `json:"counter_id,omitempty" xml:"counter_id,omitempty"`
}
