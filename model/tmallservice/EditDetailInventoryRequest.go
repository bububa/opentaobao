package tmallservice

// EditDetailInventoryRequest 
type EditDetailInventoryRequest struct {

    // 目标库存key。如果容量的时间间隔为天时，则必须为yyyy-MM-dd格式；如果容量的时间间隔为小时，则必须为yyyy-MM-dd HH:00-HH:00格式
    TargetInventoryKey   string `json:"target_inventory_key,omitempty"`

    // 不为0的整数
    Quantity   int64 `json:"quantity,omitempty"`

}
