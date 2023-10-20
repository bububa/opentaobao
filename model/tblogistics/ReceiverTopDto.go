package tblogistics

import (
	"sync"
)

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

var poolReceiverTopDto = sync.Pool{
	New: func() any {
		return new(ReceiverTopDto)
	},
}

// GetReceiverTopDto() 从对象池中获取ReceiverTopDto
func GetReceiverTopDto() *ReceiverTopDto {
	return poolReceiverTopDto.Get().(*ReceiverTopDto)
}

// ReleaseReceiverTopDto 释放ReceiverTopDto
func ReleaseReceiverTopDto(v *ReceiverTopDto) {
	v.ItemName = ""
	v.Name = ""
	v.Phone = ""
	v.Address = ""
	v.Lat = ""
	v.Lng = ""
	v.ItemValue = 0
	v.ItemQuantity = 0
	v.ItemId = 0
	poolReceiverTopDto.Put(v)
}
