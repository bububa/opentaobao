package qimen

// OrderProcessQueryResponse 
type OrderProcessQueryResponse struct {

    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty"`

    // 响应码
    Code   string `json:"code,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 订单处理流程
    OrderProcess   *OrderProcess `json:"orderProcess,omitempty"`

}
