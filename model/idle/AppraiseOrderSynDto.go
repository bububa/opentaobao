package idle

// AppraiseOrderSynDto 
type AppraiseOrderSynDto struct {
    // biz_order_id
    BizOrderId   string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    // (主状态,子状态,状态说明)示例如下： ("1", "1", "买家拍下未付款") ("2", "1", "买家拍下已付款") ("3", "1", "卖家已发货") 等，详情参考对接文档
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // 见order_status字段说明
    OrderSubStatus   string `json:"order_sub_status,omitempty" xml:"order_sub_status,omitempty"`
    // 根据订单状态不同，传递不同的内容
    Attribute   *Attribute `json:"attribute,omitempty" xml:"attribute,omitempty"`
}
