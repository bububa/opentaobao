package wms

// Itemdetailwlbwmsconsignordernotify 
type Itemdetailwlbwmsconsignordernotify struct {
    // 金额
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    // 单位
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    // 商品名称
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    // 商品价格
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    // 商品数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
