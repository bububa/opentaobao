package cityretail

// ChangeOrderResponseDTO 
type ChangeOrderResponseDTO struct {
    // 淘宝订单id
    TbOrderId   string `json:"tb_order_id,omitempty" xml:"tb_order_id,omitempty"`
    // 转仓后的新仓
    NewWarehouseCode   string `json:"new_warehouse_code,omitempty" xml:"new_warehouse_code,omitempty"`
    // 转仓前的原仓
    OriginWarehouseCode   string `json:"origin_warehouse_code,omitempty" xml:"origin_warehouse_code,omitempty"`
    // 履约单id
    FulfillOrderId   string `json:"fulfill_order_id,omitempty" xml:"fulfill_order_id,omitempty"`
    // 取货码
    PickupCode   string `json:"pickup_code,omitempty" xml:"pickup_code,omitempty"`
}
