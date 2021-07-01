package moscm

// IsvInboundRequestItemDto 
type IsvInboundRequestItemDto struct {
    // 库位编号
    LocationId   string `json:"location_id,omitempty" xml:"location_id,omitempty"`
    // 外部id
    OutId   string `json:"out_id,omitempty" xml:"out_id,omitempty"`
    // 获取或设置货物数量计量单位
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 单位
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    // 单价
    UnitPrice   string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
}
