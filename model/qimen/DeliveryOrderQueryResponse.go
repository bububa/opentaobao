package qimen

// DeliveryOrderQueryResponse 
type DeliveryOrderQueryResponse struct {

    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty"`

    // 响应码
    Code   string `json:"code,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

    // orderLines总条数
    TotalLines   int64 `json:"totalLines,omitempty"`

    // 发货单信息
    DeliveryOrder   *DeliveryOrder `json:"deliveryOrder,omitempty"`

    // 包裹信息
    Packages   []Package `json:"packages,omitempty"`

    // 单据列表
    OrderLines   []OrderLine `json:"orderLines,omitempty"`

}