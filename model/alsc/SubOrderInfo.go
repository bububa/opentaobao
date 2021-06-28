package alsc

// SubOrderInfo 
type SubOrderInfo struct {

    // 商品附加费金额 如配料，做法
    
    ItemAttachFee   int64 `json:"item_attach_fee,omitempty" xml:"item_attach_fee,omitempty"`
    

    // 商品ID
    
    OutItemId   string `json:"out_item_id,omitempty" xml:"out_item_id,omitempty"`
    

    // 商品的名称
    
    OutItemName   string `json:"out_item_name,omitempty" xml:"out_item_name,omitempty"`
    

    // 主单号
    
    OutOrderNo   string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
    

    // 商品规格id
    
    OutSkuId   string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
    

    // 商品规格名称
    
    OutSkuName   string `json:"out_sku_name,omitempty" xml:"out_sku_name,omitempty"`
    

    // 子订单号
    
    OutSubOrderNo   string `json:"out_sub_order_no,omitempty" xml:"out_sub_order_no,omitempty"`
    

    // 商品单价，如果是称重菜，等于subTotalFee
    
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    

    // 份数，如果是称重菜，值直接传1
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 商品总金额,包含配料，做法
    
    SubTotalFee   int64 `json:"sub_total_fee,omitempty" xml:"sub_total_fee,omitempty"`
    

    // 单位
    
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    

    // 重量
    
    Weight   string `json:"weight,omitempty" xml:"weight,omitempty"`
    

}
