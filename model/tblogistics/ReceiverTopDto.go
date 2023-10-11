package tblogistics

// ReceiverTopDto 结构体
type ReceiverTopDto struct {
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 联系人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 联系人电话，支持手机、座机
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 地址门牌号
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 纬度（高德）
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 经度（高德）
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 商品单价（原价）
	ItemValue int64 `json:"item_value,omitempty" xml:"item_value,omitempty"`
	// 商品数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
