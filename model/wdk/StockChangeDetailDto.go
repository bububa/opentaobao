package wdk

// StockChangeDetailDto 
type StockChangeDetailDto struct {
    // batchCode
    BatchCode   string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
    // quantity
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // itemCode
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    // reason
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    // bizOrderCode
    BizOrderCode   string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
    // cabinetCode
    CabinetCode   string `json:"cabinet_code,omitempty" xml:"cabinet_code,omitempty"`
}
