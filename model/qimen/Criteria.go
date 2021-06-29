package qimen

// Criteria 
type Criteria struct {
    // 仓库编码
    WarehouseCode   string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
    // 货主编码
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
    // 商品编码
    ItemCode   string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
    // 仓储系统商品ID
    ItemId   string `json:"itemId,omitempty" xml:"itemId,omitempty"`
    // 库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存;默认为查所有类型的库存)
    InventoryType   string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
}
