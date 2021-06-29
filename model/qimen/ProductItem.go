package qimen

// ProductItem 
type ProductItem struct {
    // erp系统商品编码
    ItemCode   string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
    // 仓储系统商品ID
    ItemId   string `json:"itemId,omitempty" xml:"itemId,omitempty"`
    // 库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;默认为ZP)
    InventoryType   string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
    // 数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 商品生产日期(YYYY-MM-DD)
    ProductDate   string `json:"productDate,omitempty" xml:"productDate,omitempty"`
    // 商品过期日期(YYYY-MM-DD)
    ExpireDate   string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
    // 生产批号
    ProduceCode   string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
    // 批次编码
    BatchCode   string `json:"batchCode,omitempty" xml:"batchCode,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 货主编码
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
    // 配比数量
    RatioQty   int64 `json:"ratioQty,omitempty" xml:"ratioQty,omitempty"`
}
