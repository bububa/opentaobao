package qimen

// EntryOrderConfirmRequest 
type EntryOrderConfirmRequest struct {

    // 入库单信息
    EntryOrder   *EntryOrder `json:"entryOrder,omitempty"`

    // 订单信息
    OrderLines   []OrderLine `json:"orderLines,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
