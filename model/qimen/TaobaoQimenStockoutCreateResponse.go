package qimen

// TaobaoQimenStockoutCreateResponse 
type TaobaoQimenStockoutCreateResponse struct {
    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`
    // 响应码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 响应信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 订单创建时间(YYYY-MM-DD HH:MM:SS)
    CreateTime   string `json:"createTime,omitempty" xml:"createTime,omitempty"`
    // 出库单仓储系统编码
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty" xml:"deliveryOrderId,omitempty"`
}
