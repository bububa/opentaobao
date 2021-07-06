package ascpchannel

// Instoragedetails 结构体
type Instoragedetails struct {
	// 库存类型:101=残次品;1=正品
	StorageType string `json:"storage_type,omitempty" xml:"storage_type,omitempty"`
	// 实际收货数量
	ReceivedQuantity int64 `json:"received_quantity,omitempty" xml:"received_quantity,omitempty"`
}
