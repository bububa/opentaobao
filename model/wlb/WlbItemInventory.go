package wlb

// WlbItemInventory 
type WlbItemInventory struct {

    // 商品id
    ItemId   int64 `json:"item_id,omitempty"`

    // 库存数量
    Quantity   int64 `json:"quantity,omitempty"`

    // 仓库编码
    StoreCode   string `json:"store_code,omitempty"`

    // 锁定库存数量
    LockQuantity   int64 `json:"lock_quantity,omitempty"`

    // SELLALBE 可销售库存<br/>DEFECTIVE 残次<br/>JISHUN 机损<br/>XIANGSHUN 箱损<br/>FREEZE 冻结库存<br/>ONROAD 在途库存
    Type   string `json:"type,omitempty"`

}
