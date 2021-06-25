package fenxiao

// InventorySum 
type InventorySum struct {

    // 商品后端ID，如果有传sc_item_code,参数可以为0
    ScItemId   int64 `json:"sc_item_id,omitempty"`

    // 商品商家编码
    ScItemCode   string `json:"sc_item_code,omitempty"`

    // 商家仓库编码
    StoreCode   string `json:"store_code,omitempty"`

    // 库存类型：<br/>1：正常 <br/>2：损坏 <br/>3：冻结 <br/>10：质押 <br/>11-20:商家自定义
    InventoryType   int64 `json:"inventory_type,omitempty"`

    // 库存类型名称
    InventoryTypeName   string `json:"inventory_type_name,omitempty"`

    // 总物理库存数量
    Quantity   int64 `json:"quantity,omitempty"`

    // 总预扣数量
    ReserveQuantity   int64 `json:"reserve_quantity,omitempty"`

    // 总占用数量
    OccupyQuantity   int64 `json:"occupy_quantity,omitempty"`

}
