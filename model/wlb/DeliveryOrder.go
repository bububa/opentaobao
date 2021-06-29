package wlb

// Deliveryorder 
type Deliveryorder struct {
    // 物流单号
    CnOrderCode   string `json:"cn_order_code,omitempty" xml:"cn_order_code,omitempty"`
    // 仓库编码
    WarehouseCode   string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
    // 创建时间
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 订单行
    OrderLines   []Orderline `json:"order_lines,omitempty" xml:"order_lines>orderline,omitempty"`
}
