package qimen

// EntryOrderCreateRequest 
type EntryOrderCreateRequest struct {

    // 入库单信息
    EntryOrder   *EntryOrder `json:"entryOrder,omitempty"`

    // 入库单详情
    OrderLines   []OrderLine `json:"orderLines,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}