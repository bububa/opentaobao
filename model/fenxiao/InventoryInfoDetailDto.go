package fenxiao

// InventoryInfoDetailDto 
type InventoryInfoDetailDto struct {

    // 占用库存数
    OccupyQuantity   int64 `json:"occupy_quantity,omitempty"`

    // 仓库物理库存数
    Quantity   int64 `json:"quantity,omitempty"`

    // 预扣库存数
    ReserveQuantity   int64 `json:"reserve_quantity,omitempty"`

    // 后端商品code
    ScItemCode   string `json:"sc_item_code,omitempty"`

    // 后端商品id
    ScItemId   int64 `json:"sc_item_id,omitempty"`

    // 仓库code
    StoreCode   string `json:"store_code,omitempty"`

    // distType
    InvStoreType   int64 `json:"inv_store_type,omitempty"`

    // subList
    SubList   []InventorySubDetailDto `json:"sub_list,omitempty"`

    // skuId
    SkuId   int64 `json:"sku_id,omitempty"`

    // 1前端商品 2供应链货品
    ItemType   int64 `json:"item_type,omitempty"`

}
