package moscm

// VirtualInventoryQueryDto 
type VirtualInventoryQueryDto struct {
    // 银泰专柜号
    CounterId   string `json:"counter_id,omitempty" xml:"counter_id,omitempty"`
    // 外部专柜号（在供应商系统中的专柜号，两个专柜号必须至少传一个，如果都传一counter_id为准）
    OutCounterId   string `json:"out_counter_id,omitempty" xml:"out_counter_id,omitempty"`
    // 外部商品id（最大列表长度：50）
    OutIds   []string `json:"out_ids,omitempty" xml:"out_ids>string,omitempty"`
}
