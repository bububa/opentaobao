package ascpchannel

// Detailorderdto 
type Detailorderdto struct {
    // 业务发生时间
    BizDate   string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
    // 实际操作子单id
    OperationDetailOrderId   string `json:"operation_detail_order_id,omitempty" xml:"operation_detail_order_id,omitempty"`
}
