package qimen

// OrderLine 
type OrderLine struct {

    // 单据行号
    OrderLineNo   string `json:"orderLineNo,omitempty"`

    // 平台交易订单编码
    OrderSourceCode   string `json:"orderSourceCode,omitempty"`

    // 平台交易子订单编码
    SubSourceCode   string `json:"subSourceCode,omitempty"`

    // 货主编码
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 商品编码
    ItemCode   string `json:"itemCode,omitempty"`

    // 商品仓储系统编码
    ItemId   string `json:"itemId,omitempty"`

    // 库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存;默认为查所有类型的库存)
    InventoryType   string `json:"inventoryType,omitempty"`

    // 商品名称
    ItemName   string `json:"itemName,omitempty"`

    // 交易平台商品编码
    ExtCode   string `json:"extCode,omitempty"`

    // 应发商品数量
    PlanQty   int64 `json:"planQty,omitempty"`

    // 实发商品数量
    ActualQty   int64 `json:"actualQty,omitempty"`

    // 批次编号
    BatchCode   string `json:"batchCode,omitempty"`

    // 生产日期(YYYY-MM-DD)
    ProductDate   string `json:"productDate,omitempty"`

    // 过期日期(YYYY-MM-DD)
    ExpireDate   string `json:"expireDate,omitempty"`

    // 生产批号
    ProduceCode   string `json:"produceCode,omitempty"`

    // 批次列表
    Batchs   []Batch `json:"batchs,omitempty"`

    // 商品的二维码(类似电子产品的SN码;用来进行商品的溯源;多个二维码之间用分号;隔开)
    QrCode   string `json:"qrCode,omitempty"`

    // 仓库拆单子发货单号
    SubDeliveryOrderId   string `json:"subDeliveryOrderId,omitempty"`

    // snList
    SnList   *SnList `json:"snList,omitempty"`

    // 供应商编码
    SupplierCode   string `json:"supplierCode,omitempty"`

    // 供应商名称
    SupplierName   string `json:"supplierName,omitempty"`

    // 交易平台订单
    SourceOrderCode   string `json:"sourceOrderCode,omitempty"`

    // 交易平台子订单编码
    SubSourceOrderCode   string `json:"subSourceOrderCode,omitempty"`

    // 支付平台交易号(淘系订单传支付宝交易号)
    PayNo   string `json:"payNo,omitempty"`

    // 零售价(零售价=实际成交价+单件商品折扣金额)
    RetailPrice   string `json:"retailPrice,omitempty"`

    // 实际成交价
    ActualPrice   string `json:"actualPrice,omitempty"`

    // 单件商品折扣金额
    DiscountAmount   string `json:"discountAmount,omitempty"`

    // 数量
    Quantity   string `json:"quantity,omitempty"`

    // 货品sn编码
    SnCode   string `json:"snCode,omitempty"`

    // 外部业务编码(消息ID;用于去重;当单据需要分批次发送时使用)
    OutBizCode   string `json:"outBizCode,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 单位
    Unit   string `json:"unit,omitempty"`

    // 商品属性
    SkuProperty   string `json:"skuProperty,omitempty"`

    // 采购价
    PurchasePrice   string `json:"purchasePrice,omitempty"`

    // 出库单号, string (50) , 必填
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty"`

    // 奇门仓储字段
    ProductCode   string `json:"productCode,omitempty"`

    // 奇门仓储字段
    StockInQty   string `json:"stockInQty,omitempty"`

    // 奇门仓储字段
    StockOutQty   string `json:"stockOutQty,omitempty"`

    // 奇门仓储字段
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 奇门仓储字段
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty"`

    // 奇门仓储字段
    Status   string `json:"status,omitempty"`

    // 奇门仓储字段
    TaobaoItemCode   string `json:"taobaoItemCode,omitempty"`

    // 奇门仓储字段
    DiscountPrice   string `json:"discountPrice,omitempty"`

    // 奇门仓储字段
    Color   string `json:"color,omitempty"`

    // 奇门仓储字段
    Size   string `json:"size,omitempty"`

    // 奇门仓储字段
    StandardPrice   string `json:"standardPrice,omitempty"`

    // 奇门仓储字段
    ReferencePrice   string `json:"referencePrice,omitempty"`

    // 奇门仓储字段
    Discount   string `json:"discount,omitempty"`

    // 奇门仓储字段
    StandardAmount   string `json:"standardAmount,omitempty"`

    // 奇门仓储字段
    SettlementAmount   string `json:"settlementAmount,omitempty"`

    // 奇门仓储字段
    LocationCode   string `json:"locationCode,omitempty"`

    // 奇门仓储字段
    Amount   string `json:"amount,omitempty"`

    // 奇门仓储字段
    MoveOutLocation   string `json:"moveOutLocation,omitempty"`

    // 奇门仓储字段
    MoveInLocation   string `json:"moveInLocation,omitempty"`

    // 奇门仓储字段
    ExceptionQty   string `json:"exceptionQty,omitempty"`

    // 用字符串格式来表示订单标记列表(比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL等;中间用“^”来隔开订单标记list (所 有字母全部大写) VISIT=上门；SELLER_AFFORD=是否卖家承担运费(默认是)SYNC_RETURN_BILL=同时退回发票)
    OrderFlag   string `json:"orderFlag,omitempty"`

    // 退货原因
    ReturnReason   string `json:"returnReason,omitempty"`

    // 交易平台商品编码
    PlatformCode   string `json:"platformCode,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
