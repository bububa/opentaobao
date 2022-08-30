package ascpchannel

// WaybillQueryPackageItem 结构体
type WaybillQueryPackageItem struct {
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品code
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 商品数量
	ItemCount string `json:"item_count,omitempty" xml:"item_count,omitempty"`
}
