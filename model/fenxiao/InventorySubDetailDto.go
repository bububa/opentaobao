package fenxiao

// InventorySubDetailDto 
/* model for simplify = false
type InventorySubDetailDto struct {

    // 预扣库存数
    
    ReserveQuantity   int64 `json:"reserve_quantity,omitempty"`
    

    // 仓库物理库存数
    
    Quantity   int64 `json:"quantity,omitempty"`
    

    // 占用库存数
    
    OccupyQuantity   int64 `json:"occupy_quantity,omitempty"`
    

    // ONLINE_INVENTORY:线上可售卖库存。SHARE_INVENTORY：线下独享库存，门店自提可用
    
    InvBizCode   string `json:"inv_biz_code,omitempty"`
    

}
*/

// InventorySubDetailDto 
type InventorySubDetailDto struct {

    // 预扣库存数
    ReserveQuantity   int64 `json:"reserve_quantity,omitempty"`

    // 仓库物理库存数
    Quantity   int64 `json:"quantity,omitempty"`

    // 占用库存数
    OccupyQuantity   int64 `json:"occupy_quantity,omitempty"`

    // ONLINE_INVENTORY:线上可售卖库存。SHARE_INVENTORY：线下独享库存，门店自提可用
    InvBizCode   string `json:"inv_biz_code,omitempty"`

}
