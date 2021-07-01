package rhino

// ClothingSkuDto 结构体
type ClothingSkuDto struct {
	// 发货数量item_count
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 成衣物料名称ItemName
	StyleName string `json:"style_name,omitempty" xml:"style_name,omitempty"`
	// 成衣sku编码Item
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 成衣rfid编码
	Rfid string `json:"rfid,omitempty" xml:"rfid,omitempty"`
}
