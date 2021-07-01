package qimen

// TaobaoQimenOrderQueryResponse 
type TaobaoQimenOrderQueryResponse struct {
    // success|failure，必填
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`
    // 响应码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 响应信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 订单列表
    OrderLines   *OrderLines `json:"orderLines,omitempty" xml:"orderLines,omitempty"`
}
