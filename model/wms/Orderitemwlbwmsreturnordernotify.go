package wms

// Orderitemwlbwmsreturnordernotify 
type Orderitemwlbwmsreturnordernotify struct {

    // 平台交易订单编码,淘系交易请传入交易单号
    OrderItemId   string `json:"order_item_id,omitempty"`

    // 平台交易订单编码,淘系交易请传入交易单号
    OrderSourceCode   string `json:"order_source_code,omitempty"`

    // 平台交易子订单编码，交易单号传入，子交易编号必填
    SubSourceCode   string `json:"sub_source_code,omitempty"`

    // 后端商品ID
    ItemId   string `json:"item_id,omitempty"`

    // 商品数量
    ItemQuantity   int64 `json:"item_quantity,omitempty"`

    // 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
    ExtendFields   string `json:"extend_fields,omitempty"`

    // 商品名称
    ItemName   string `json:"item_name,omitempty"`

}
