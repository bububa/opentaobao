package qimen

// StockOutCreateRequest 
type StockOutCreateRequest struct {

    // 出库单信息
    DeliveryOrder   *DeliveryOrder `json:"deliveryOrder,omitempty"`

    // 单据信息
    OrderLines   []OrderLine `json:"orderLines,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
