package game

// AppleTopOldSignNotifyDo 
type AppleTopOldSignNotifyDo struct {

    // 电子卡卡号
    
    CardNo   string `json:"card_no,omitempty" xml:"card_no,omitempty"`
    

    // 商品编号
    
    GoodsId   string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
    

    // 商户订单号
    
    OrderNo   string `json:"order_no,omitempty" xml:"order_no,omitempty"`
    

    // 附加信息，后续可以扩展
    
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    

    // 面值，单位分
    
    FacePrice   string `json:"face_price,omitempty" xml:"face_price,omitempty"`
    

    // 网关订单号
    
    GatewayOrderNo   string `json:"gateway_order_no,omitempty" xml:"gateway_order_no,omitempty"`
    

}
