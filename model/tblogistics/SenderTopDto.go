package tblogistics

import (
	"sync"
)

// SenderTopDto 结构体
type SenderTopDto struct {
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
}

var poolSenderTopDto = sync.Pool{
	New: func() any {
		return new(SenderTopDto)
	},
}

// GetSenderTopDto() 从对象池中获取SenderTopDto
func GetSenderTopDto() *SenderTopDto {
	return poolSenderTopDto.Get().(*SenderTopDto)
}

// ReleaseSenderTopDto 释放SenderTopDto
func ReleaseSenderTopDto(v *SenderTopDto) {
	v.Name = ""
	v.Phone = ""
	v.Address = ""
	v.Lat = ""
	v.Lng = ""
	poolSenderTopDto.Put(v)
}
