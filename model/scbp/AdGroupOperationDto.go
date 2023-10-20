package scbp

import (
	"sync"
)

// AdGroupOperationDto 结构体
type AdGroupOperationDto struct {
	// key
	SettingKey string `json:"setting_key,omitempty" xml:"setting_key,omitempty"`
	// value
	SettingValue string `json:"setting_value,omitempty" xml:"setting_value,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 线上状态
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 产品id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolAdGroupOperationDto = sync.Pool{
	New: func() any {
		return new(AdGroupOperationDto)
	},
}

// GetAdGroupOperationDto() 从对象池中获取AdGroupOperationDto
func GetAdGroupOperationDto() *AdGroupOperationDto {
	return poolAdGroupOperationDto.Get().(*AdGroupOperationDto)
}

// ReleaseAdGroupOperationDto 释放AdGroupOperationDto
func ReleaseAdGroupOperationDto(v *AdGroupOperationDto) {
	v.SettingKey = ""
	v.SettingValue = ""
	v.ProductId = 0
	v.OnlineStatus = 0
	v.Id = 0
	poolAdGroupOperationDto.Put(v)
}
