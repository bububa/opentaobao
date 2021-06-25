package wdk

// WarehouseSkuDo 
type WarehouseSkuDo struct {

    // 商品条码
    Barcodes   []WarehouseSkuBarcodeDo `json:"barcodes,omitempty"`

    // 品牌名称
    BrandName   string `json:"brand_name,omitempty"`

    // 含量
    Content   string `json:"content,omitempty"`

    // 配货规格
    DeliverySpec   string `json:"delivery_spec,omitempty"`

    // 配货单位
    DeliveryUnit   string `json:"delivery_unit,omitempty"`

    // 配送方式，1-统配、2-直配、3-越库
    DeliveryWay   string `json:"delivery_way,omitempty"`

    // 禁收时限
    ForbidReceiveDays   int64 `json:"forbid_receive_days,omitempty"`

    // 禁售时限
    ForbidSalesDays   int64 `json:"forbid_sales_days,omitempty"`

    // 新建时间
    GmtCreateTime   string `json:"gmt_create_time,omitempty"`

    // 是否进口商品
    ImportFlag   bool `json:"import_flag,omitempty"`

    // 进项税率
    InputTaxRate   string `json:"input_tax_rate,omitempty"`

    // 库存单位
    InventoryUnit   string `json:"inventory_unit,omitempty"`

    // 商品状态，A-正常、T-暂时停采、C-淘汰出清、R-清退、L-季节性商品休眠、D-删除封挡、E-停售(紧急下架)、U-未启用（只是建档，还未进货）
    LifeStatus   string `json:"life_status,omitempty"`

    // 商品类目
    MerchantCatId   int64 `json:"merchant_cat_id,omitempty"`

    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 超收比例
    OverloadRate   string `json:"overload_rate,omitempty"`

    // 保质期天数
    Period   int64 `json:"period,omitempty"`

    // 厂商名称
    ProducerName   string `json:"producer_name,omitempty"`

    // 产地，多个产地使用逗号分割
    ProducerPlace   string `json:"producer_place,omitempty"`

    // 采购规格
    PurchaseSpec   string `json:"purchase_spec,omitempty"`

    // 采购单位
    PurchaseUnit   string `json:"purchase_unit,omitempty"`

    // 简称
    ShortTitle   string `json:"short_title,omitempty"`

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 商品名称
    SkuName   string `json:"sku_name,omitempty"`

    // 存储方式，241-常温、242-冷藏、243-冷冻、635-热链、636-室温、637-鲜活
    Storage   string `json:"storage,omitempty"`

    // 当前供应商编码
    SupplierNo   string `json:"supplier_no,omitempty"`

    // 进项税率
    TaxRate   string `json:"tax_rate,omitempty"`

    // 仓编码
    WarehouseCode   string `json:"warehouse_code,omitempty"`

    // 保质期预警天数
    WarnDays   int64 `json:"warn_days,omitempty"`

    // 是否称重商品
    WeightFlag   bool `json:"weight_flag,omitempty"`

}
