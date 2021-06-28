package trade

// OrderInfoDto 
type OrderInfoDto struct {

    // 订单状态描述
    
    OrderStatusDesc   string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
    

    // 店铺名称
    
    ShopName   string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
    

    // 店铺url
    
    ShopUrl   string `json:"shop_url,omitempty" xml:"shop_url,omitempty"`
    

    // 订单标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 订单总数
    
    TotalCount   string `json:"total_count,omitempty" xml:"total_count,omitempty"`
    

    // 订单总数说明
    
    TotalCountDesc   string `json:"total_count_desc,omitempty" xml:"total_count_desc,omitempty"`
    

    // 支付金额
    
    PayAmount   string `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
    

    // 商品摘要
    
    ItemDigests   []ItemDigestDto `json:"item_digests,omitempty" xml:"item_digests,omitempty"`
    

    // 订单详情url
    
    OrderDetailUrl   string `json:"order_detail_url,omitempty" xml:"order_detail_url,omitempty"`
    

    // 支付金额描述
    
    PayInfo   string `json:"pay_info,omitempty" xml:"pay_info,omitempty"`
    

}
