package fenxiao

// InventoryQueryForStoreRequest 
/* model for simplify = false
type InventoryQueryForStoreRequest struct {

    // 后端商品code， sc_item_code或 sc_item_id需传入其中之一
    
    ScItemCode   string `json:"sc_item_code,omitempty"`
    

    // 仓库code
    
    StoreCode   string `json:"store_code,omitempty"`
    

    // 后端商品id， sc_item_code或 sc_item_id需传入其中之一
    
    ScItemId   int64 `json:"sc_item_id,omitempty"`
    

    // 实体类型  2：仓库类型， 6：门店类型
    
    InvStoreType   int64 `json:"inv_store_type,omitempty"`
    

    // 货品填0即可，前端商品填skuId
    
    SkuId   int64 `json:"sku_id,omitempty"`
    

    // 1前端商品 2供应链货品
    
    ItemType   int64 `json:"item_type,omitempty"`
    

}
*/

// InventoryQueryForStoreRequest 
type InventoryQueryForStoreRequest struct {

    // 后端商品code， sc_item_code或 sc_item_id需传入其中之一
    ScItemCode   string `json:"sc_item_code,omitempty"`

    // 仓库code
    StoreCode   string `json:"store_code,omitempty"`

    // 后端商品id， sc_item_code或 sc_item_id需传入其中之一
    ScItemId   int64 `json:"sc_item_id,omitempty"`

    // 实体类型  2：仓库类型， 6：门店类型
    InvStoreType   int64 `json:"inv_store_type,omitempty"`

    // 货品填0即可，前端商品填skuId
    SkuId   int64 `json:"sku_id,omitempty"`

    // 1前端商品 2供应链货品
    ItemType   int64 `json:"item_type,omitempty"`

}
