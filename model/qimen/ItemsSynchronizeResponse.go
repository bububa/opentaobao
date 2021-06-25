package qimen

// ItemsSynchronizeResponse 
type ItemsSynchronizeResponse struct {

    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty"`

    // 响应码
    Code   string `json:"code,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 商品同步批量接口中单商品信息
    Items   []BatchItemSynItem `json:"items,omitempty"`

}
