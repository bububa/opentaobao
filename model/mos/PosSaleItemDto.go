package mos

// PosSaleItemDto 
/* model for simplify = false
type PosSaleItemDto struct {

    // 门店号
    
    MallNo   string `json:"mall_no,omitempty"`
    

    // 专柜号
    
    ShopNo   string `json:"shop_no,omitempty"`
    

    // 开票单号（）
    
    SaleTicketNo   string `json:"sale_ticket_no,omitempty"`
    

    // 商品类型：  N：老单品；Y：喵货；C：云单品；M：零售+
    
    ItemType   string `json:"item_type,omitempty"`
    

    // 商品行号
    
    GoodsLineNo   int64 `json:"goods_line_no,omitempty"`
    

    // 商品ID
    
    GoodsId   string `json:"goods_id,omitempty"`
    

    // 唯一码
    
    UniqueCode   string `json:"unique_code,omitempty"`
    

    // 商品名称
    
    GoodsName   string `json:"goods_name,omitempty"`
    

    // 数量，可小数点
    
    Quantity   string `json:"quantity,omitempty"`
    

    // 商品原价，单位：分
    
    OriPrice   int64 `json:"ori_price,omitempty"`
    

    // 折扣金额，单位：分
    
    DiscountAmount   int64 `json:"discount_amount,omitempty"`
    

    // 应付金额，单位：分
    
    PayAmount   int64 `json:"pay_amount,omitempty"`
    

    // 结算码（退款的时候必须有）
    
    SettleCode   string `json:"settle_code,omitempty"`
    

    // 扩展参数
    
    ExtendParams   string `json:"extend_params,omitempty"`
    

}
*/

// PosSaleItemDto 
type PosSaleItemDto struct {

    // 门店号
    MallNo   string `json:"mall_no,omitempty"`

    // 专柜号
    ShopNo   string `json:"shop_no,omitempty"`

    // 开票单号（）
    SaleTicketNo   string `json:"sale_ticket_no,omitempty"`

    // 商品类型：  N：老单品；Y：喵货；C：云单品；M：零售+
    ItemType   string `json:"item_type,omitempty"`

    // 商品行号
    GoodsLineNo   int64 `json:"goods_line_no,omitempty"`

    // 商品ID
    GoodsId   string `json:"goods_id,omitempty"`

    // 唯一码
    UniqueCode   string `json:"unique_code,omitempty"`

    // 商品名称
    GoodsName   string `json:"goods_name,omitempty"`

    // 数量，可小数点
    Quantity   string `json:"quantity,omitempty"`

    // 商品原价，单位：分
    OriPrice   int64 `json:"ori_price,omitempty"`

    // 折扣金额，单位：分
    DiscountAmount   int64 `json:"discount_amount,omitempty"`

    // 应付金额，单位：分
    PayAmount   int64 `json:"pay_amount,omitempty"`

    // 结算码（退款的时候必须有）
    SettleCode   string `json:"settle_code,omitempty"`

    // 扩展参数
    ExtendParams   string `json:"extend_params,omitempty"`

}
