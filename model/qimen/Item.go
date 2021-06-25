package qimen

// Item 
type Item struct {

    // 商品编码
    ItemCode   string `json:"itemCode,omitempty"`

    // 后端商品编码
    ItemId   string `json:"itemId,omitempty"`

    // 组合商品中的该商品个数
    Quantity   int64 `json:"quantity,omitempty"`

    // ownerCode
    OwnerCode   string `json:"ownerCode,omitempty"`

    // test
    BrandName   string `json:"brandName,omitempty"`

    // test
    Sn   string `json:"sn,omitempty"`

    // test
    IsSNMgmt   string `json:"isSNMgmt,omitempty"`

    // test
    BarCode   string `json:"barCode,omitempty"`

    // test
    Color   string `json:"color,omitempty"`

    // test
    Size   string `json:"size,omitempty"`

    // test
    Length   string `json:"length,omitempty"`

    // test
    Width   string `json:"width,omitempty"`

    // test
    Height   string `json:"height,omitempty"`

    // test
    Volume   string `json:"volume,omitempty"`

    // test
    GrossWeight   string `json:"grossWeight,omitempty"`

    // test
    NetWeight   string `json:"netWeight,omitempty"`

    // test
    TareWeight   string `json:"tareWeight,omitempty"`

    // test
    SafetyStock   string `json:"safetyStock,omitempty"`

    // test
    StockUnit   string `json:"stockUnit,omitempty"`

    // test
    StockStatus   string `json:"stockStatus,omitempty"`

    // test
    ProductDate   string `json:"productDate,omitempty"`

    // test
    ExpireDate   string `json:"expireDate,omitempty"`

    // test
    IsShelfLifeMgmt   string `json:"isShelfLifeMgmt,omitempty"`

    // test
    ShelfLife   string `json:"shelfLife,omitempty"`

    // test
    RejectLifecycle   string `json:"rejectLifecycle,omitempty"`

    // test
    LockupLifecycle   string `json:"lockupLifecycle,omitempty"`

    // test
    AdventLifecycle   string `json:"adventLifecycle,omitempty"`

    // test
    BatchCode   string `json:"batchCode,omitempty"`

    // test
    BatchRemark   string `json:"batchRemark,omitempty"`

    // test
    IsBatchMgmt   string `json:"isBatchMgmt,omitempty"`

    // test
    PackCode   string `json:"packCode,omitempty"`

    // test
    Pcs   string `json:"pcs,omitempty"`

    // test
    OriginAddress   string `json:"originAddress,omitempty"`

    // test
    ApprovalNumber   string `json:"approvalNumber,omitempty"`

    // test
    IsFragile   string `json:"isFragile,omitempty"`

    // test
    IsHazardous   string `json:"isHazardous,omitempty"`

    // test
    PricingCategory   string `json:"pricingCategory,omitempty"`

    // test
    IsSku   string `json:"isSku,omitempty"`

    // test
    PackageMaterial   string `json:"packageMaterial,omitempty"`

    // test
    Price   string `json:"price,omitempty"`

    // test
    IsAreaSale   string `json:"isAreaSale,omitempty"`

    // test
    NormalQty   string `json:"normalQty,omitempty"`

    // test
    DefectiveQty   string `json:"defectiveQty,omitempty"`

    // test
    ReceiveQty   string `json:"receiveQty,omitempty"`

    // test
    ExCode   string `json:"exCode,omitempty"`

    // test
    DiscountPrice   string `json:"discountPrice,omitempty"`

    // test
    InventoryType   string `json:"inventoryType,omitempty"`

    // test
    PlanQty   string `json:"planQty,omitempty"`

    // test
    SourceOrderCode   string `json:"sourceOrderCode,omitempty"`

    // test
    SubSourceOrderCode   string `json:"subSourceOrderCode,omitempty"`

    // test
    ProduceCode   string `json:"produceCode,omitempty"`

    // test
    OrderLineNo   string `json:"orderLineNo,omitempty"`

    // test
    ActualQty   string `json:"actualQty,omitempty"`

    // test
    Amount   string `json:"amount,omitempty"`

    // test
    Unit   string `json:"unit,omitempty"`

    // test
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // test
    LockQuantity   string `json:"lockQuantity,omitempty"`

    // test
    OrderCode   string `json:"orderCode,omitempty"`

    // test
    OrderType   string `json:"orderType,omitempty"`

    // test
    OutBizCode   string `json:"outBizCode,omitempty"`

    // test
    ProductCode   string `json:"productCode,omitempty"`

    // test
    PaperQty   string `json:"paperQty,omitempty"`

    // test
    DiffQuantity   string `json:"diffQuantity,omitempty"`

    // test
    ExtCode   string `json:"extCode,omitempty"`

    // test
    LackQty   string `json:"lackQty,omitempty"`

    // test
    Reason   string `json:"reason,omitempty"`

    // test
    SnCode   string `json:"snCode,omitempty"`

    // test
    GoodsCode   string `json:"goodsCode,omitempty"`

    // test
    StandardPrice   string `json:"standardPrice,omitempty"`

    // test
    ReferencePrice   string `json:"referencePrice,omitempty"`

    // test
    Discount   string `json:"discount,omitempty"`

    // test
    ActualAmount   string `json:"actualAmount,omitempty"`

    // test
    PriceAdjustment   *PriceAdjustment `json:"priceAdjustment,omitempty"`

    // test
    LatestUpdateTime   string `json:"latestUpdateTime,omitempty"`

    // test
    ChangeTime   string `json:"changeTime,omitempty"`

    // test
    TempRequirement   string `json:"tempRequirement,omitempty"`

    // test
    ChannelCode   string `json:"channelCode,omitempty"`

    // test
    OriginCode   string `json:"originCode,omitempty"`

    // test
    Batchs   []Batch `json:"batchs,omitempty"`

    // 商品名称
    ItemName   string `json:"itemName,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 库存商品总量
    TotalQty   int64 `json:"totalQty,omitempty"`

    // 商品简称
    ShortName   string `json:"shortName,omitempty"`

    // 英文名
    EnglishName   string `json:"englishName,omitempty"`

    // 商品属性(如红色;XXL)
    SkuProperty   string `json:"skuProperty,omitempty"`

    // 渠道中的商品标题
    Title   string `json:"title,omitempty"`

    // 商品类别ID
    CategoryId   string `json:"categoryId,omitempty"`

    // 商品类别名称
    CategoryName   string `json:"categoryName,omitempty"`

    // 商品类型(ZC=正常商品;FX=分销商品;ZH=组合商品;ZP=赠品;BC=包材;HC=耗材;FL=辅料;XN=虚拟品;FS=附属品;CC=残次品;OTHER=其它;只传英文编码)
    ItemType   string `json:"itemType,omitempty"`

    // 吊牌价
    TagPrice   string `json:"tagPrice,omitempty"`

    // 零售价
    RetailPrice   string `json:"retailPrice,omitempty"`

    // 成本价
    CostPrice   string `json:"costPrice,omitempty"`

    // 采购价
    PurchasePrice   string `json:"purchasePrice,omitempty"`

    // 季节编码
    SeasonCode   string `json:"seasonCode,omitempty"`

    // 季节名称
    SeasonName   string `json:"seasonName,omitempty"`

    // 品牌代码
    BrandCode   string `json:"brandCode,omitempty"`

    // 创建时间(YYYY-MM-DD HH:MM:SS)
    CreateTime   string `json:"createTime,omitempty"`

    // 更新时间(YYYY-MM-DD HH:MM:SS)
    UpdateTime   string `json:"updateTime,omitempty"`

    // 是否有效(Y/N ;默认为N)
    IsValid   string `json:"isValid,omitempty"`

    // 销售配送方式（0=自配|1=菜鸟）
    LogisticsType   string `json:"logisticsType,omitempty"`

    // 是否液体, Y/N,  (默认为N)
    IsLiquid   string `json:"isLiquid,omitempty"`

    // 签收状态
    Status   string `json:"status,omitempty"`

    // temp
    SupplierCode   string `json:"supplierCode,omitempty"`

    // isLocked
    IsLocked   string `json:"isLocked,omitempty"`

    // 货品编码,HZ1234,string(50),,
    ScItemCode   string `json:"scItemCode,omitempty"`

    // 计划调拨数量,Item1234,string(50),,
    PlanCount   string `json:"planCount,omitempty"`

    // 实际出库数量,Item1234,string(50),,
    OutCount   string `json:"outCount,omitempty"`

    // 实际入库数量,Item1234,string(50),,
    InCount   string `json:"inCount,omitempty"`

}
