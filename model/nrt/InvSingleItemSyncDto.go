package nrt

// InvSingleItemSyncDto 
type InvSingleItemSyncDto struct {
    // 商品条形码,数字型，作为外部商品唯一外键优先选用该字段，唯一性由商家保证，如果没有请联系阿里小二。bar_code和outer_item_id必须传入其中一个
    BarCode   string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
    // 仓编码
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    // 业务码
    BizCode   string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
    // 商家id
    SellerId   int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
    // 如果商家没有商品条形码,用该字段，唯一性由商家保证。
    OuterItemId   string `json:"outer_item_id,omitempty" xml:"outer_item_id,omitempty"`
    // 可用销售库存
    SellableQuantity   int64 `json:"sellable_quantity,omitempty" xml:"sellable_quantity,omitempty"`
    // 库存类型，warehouse-仓库存；store-门店库存，不传默认仓库存
    InvType   string `json:"inv_type,omitempty" xml:"inv_type,omitempty"`
    // 外部操作单号，用于幂等控制
    OuterOrderNo   string `json:"outer_order_no,omitempty" xml:"outer_order_no,omitempty"`
    // 变化数量，整数表示增加数量，负数表示减少数量
    ChangeQuantity   int64 `json:"change_quantity,omitempty" xml:"change_quantity,omitempty"`
}
