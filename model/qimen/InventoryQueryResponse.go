package qimen

// InventoryQueryResponse 
type InventoryQueryResponse struct {

    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty"`

    // 响应码
    Code   string `json:"code,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 商品的库存信息列表
    Items   []Item `json:"items,omitempty"`

}
