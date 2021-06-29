package qimen

// Batch 
type Batch struct {
    // test
    BatchCode   string `json:"batchCode,omitempty" xml:"batchCode,omitempty"`
    // test
    ProductDate   string `json:"productDate,omitempty" xml:"productDate,omitempty"`
    // test
    ExpireDate   string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
    // test
    ProduceCode   string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
    // test
    InventoryType   string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
    // test
    ActualQty   string `json:"actualQty,omitempty" xml:"actualQty,omitempty"`
    // test
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 货品sn编码
    SnCode   string `json:"snCode,omitempty" xml:"snCode,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
}
