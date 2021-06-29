package wdk

// ReceiptSubOrderDO 
type ReceiptSubOrderDO struct {
    // 成交金额
    DealAmt   int64 `json:"deal_amt,omitempty" xml:"deal_amt,omitempty"`
    // 成交单价
    DealPrice   int64 `json:"deal_price,omitempty" xml:"deal_price,omitempty"`
    // 序号
    Index   int64 `json:"index,omitempty" xml:"index,omitempty"`
    // 商品条码
    ItemBarcode   string `json:"item_barcode,omitempty" xml:"item_barcode,omitempty"`
    // 商品编码
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    // 商品名称
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    // 会员优惠
    MemberDiscount   int64 `json:"member_discount,omitempty" xml:"member_discount,omitempty"`
    // 商品原售价
    OriginalPrice   int64 `json:"original_price,omitempty" xml:"original_price,omitempty"`
    // 款机号
    PosNo   string `json:"pos_no,omitempty" xml:"pos_no,omitempty"`
    // 子订单优惠总金额
    PromotionDiscount   int64 `json:"promotion_discount,omitempty" xml:"promotion_discount,omitempty"`
    // 数量
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 扫描条码
    ScanBarcode   string `json:"scan_barcode,omitempty" xml:"scan_barcode,omitempty"`
    // 售价金额
    SellingPrice   int64 `json:"selling_price,omitempty" xml:"selling_price,omitempty"`
    // 流水号
    SerialNum   string `json:"serial_num,omitempty" xml:"serial_num,omitempty"`
    // 门店号
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 临时折扣
    TemporaryDiscount   int64 `json:"temporary_discount,omitempty" xml:"temporary_discount,omitempty"`
    // 单位
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
}
