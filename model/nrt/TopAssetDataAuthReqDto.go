package nrt

import (
	"sync"
)

// TopAssetDataAuthReqDto 结构体
type TopAssetDataAuthReqDto struct {
	// 员工手机号
	PhoneList []string `json:"phone_list,omitempty" xml:"phone_list>string,omitempty"`
	// 同城站ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolTopAssetDataAuthReqDto = sync.Pool{
	New: func() any {
		return new(TopAssetDataAuthReqDto)
	},
}

// GetTopAssetDataAuthReqDto() 从对象池中获取TopAssetDataAuthReqDto
func GetTopAssetDataAuthReqDto() *TopAssetDataAuthReqDto {
	return poolTopAssetDataAuthReqDto.Get().(*TopAssetDataAuthReqDto)
}

// ReleaseTopAssetDataAuthReqDto 释放TopAssetDataAuthReqDto
func ReleaseTopAssetDataAuthReqDto(v *TopAssetDataAuthReqDto) {
	v.PhoneList = v.PhoneList[:0]
	v.StoreId = 0
	poolTopAssetDataAuthReqDto.Put(v)
}
