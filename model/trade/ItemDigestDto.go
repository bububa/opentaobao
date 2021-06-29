package trade

// ItemDigestDTO 
type ItemDigestDTO struct {
    // 图片url
    ImgUrl   string `json:"img_url,omitempty" xml:"img_url,omitempty"`
    // 商品名称
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    // 商品数量
    ItemQuantity   string `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
    // 商品金额
    ItemPrice   string `json:"item_price,omitempty" xml:"item_price,omitempty"`
    // 商品订单状态描述
    ItemOrderStatusDesc   string `json:"item_order_status_desc,omitempty" xml:"item_order_status_desc,omitempty"`
}
