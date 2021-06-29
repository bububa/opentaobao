package icbulogistics

// OrderDetailDTO 
type OrderDetailDTO struct {
    // 仓库信息
    Warehouse   *WarehouseDTO `json:"warehouse,omitempty" xml:"warehouse,omitempty"`
    // 条码Base64
    BarCode   string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
    // 物流订单号
    OrderNumber   string `json:"order_number,omitempty" xml:"order_number,omitempty"`
    // 上门揽收信息
    PickupInfo   *PickupInfoDTO `json:"pickup_info,omitempty" xml:"pickup_info,omitempty"`
}
