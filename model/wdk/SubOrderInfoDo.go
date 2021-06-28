package wdk

// SubOrderInfoDo 
/* model for simplify = false
type SubOrderInfoDo struct {

    // 单位
    
    Unit   string `json:"unit,omitempty"`
    

    // 扫描条码
    
    ScanBarcode   string `json:"scan_barcode,omitempty"`
    

    // 临时折扣
    
    TemporaryDiscount   int64 `json:"temporary_discount,omitempty"`
    

    // 子订单优惠总金额
    
    PromotionDiscount   int64 `json:"promotion_discount,omitempty"`
    

    // 会员优惠
    
    MemberDiscount   int64 `json:"member_discount,omitempty"`
    

    // 成交金额（填写会校验）
    
    DealAmt   int64 `json:"deal_amt,omitempty"`
    

    // 成交单价
    
    DealPrice   int64 `json:"deal_price,omitempty"`
    

    // 售价金额
    
    SellingPrice   int64 `json:"selling_price,omitempty"`
    

    // 商品原售价
    
    OriginalPrice   int64 `json:"original_price,omitempty"`
    

    // 数量
    
    Quantity   string `json:"quantity,omitempty"`
    

    // 商品名称
    
    ItemName   string `json:"item_name,omitempty"`
    

    // 商品条码
    
    ItemBarcode   string `json:"item_barcode,omitempty"`
    

    // 商品编码
    
    ItemCode   string `json:"item_code,omitempty"`
    

    // 序号
    
    Index   int64 `json:"index,omitempty"`
    

    // 子订单流水号
    
    SerialNum   string `json:"serial_num,omitempty"`
    

    // 款机号
    
    PosNo   string `json:"pos_no,omitempty"`
    

    // 渠道店id
    
    StoreId   string `json:"store_id,omitempty"`
    

}
*/

// SubOrderInfoDo 
type SubOrderInfoDo struct {

    // 单位
    Unit   string `json:"unit,omitempty"`

    // 扫描条码
    ScanBarcode   string `json:"scan_barcode,omitempty"`

    // 临时折扣
    TemporaryDiscount   int64 `json:"temporary_discount,omitempty"`

    // 子订单优惠总金额
    PromotionDiscount   int64 `json:"promotion_discount,omitempty"`

    // 会员优惠
    MemberDiscount   int64 `json:"member_discount,omitempty"`

    // 成交金额（填写会校验）
    DealAmt   int64 `json:"deal_amt,omitempty"`

    // 成交单价
    DealPrice   int64 `json:"deal_price,omitempty"`

    // 售价金额
    SellingPrice   int64 `json:"selling_price,omitempty"`

    // 商品原售价
    OriginalPrice   int64 `json:"original_price,omitempty"`

    // 数量
    Quantity   string `json:"quantity,omitempty"`

    // 商品名称
    ItemName   string `json:"item_name,omitempty"`

    // 商品条码
    ItemBarcode   string `json:"item_barcode,omitempty"`

    // 商品编码
    ItemCode   string `json:"item_code,omitempty"`

    // 序号
    Index   int64 `json:"index,omitempty"`

    // 子订单流水号
    SerialNum   string `json:"serial_num,omitempty"`

    // 款机号
    PosNo   string `json:"pos_no,omitempty"`

    // 渠道店id
    StoreId   string `json:"store_id,omitempty"`

}
