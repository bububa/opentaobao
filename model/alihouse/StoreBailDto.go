package alihouse

import (
	"sync"
)

// StoreBailDto 结构体
type StoreBailDto struct {
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 测试标
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 版本号
	EtcVersion int64 `json:"etc_version,omitempty" xml:"etc_version,omitempty"`
	// 保证金余额
	Bail int64 `json:"bail,omitempty" xml:"bail,omitempty"`
	// 行业类型
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 类目编码
	CategoryCode int64 `json:"category_code,omitempty" xml:"category_code,omitempty"`
}

var poolStoreBailDto = sync.Pool{
	New: func() any {
		return new(StoreBailDto)
	},
}

// GetStoreBailDto() 从对象池中获取StoreBailDto
func GetStoreBailDto() *StoreBailDto {
	return poolStoreBailDto.Get().(*StoreBailDto)
}

// ReleaseStoreBailDto 释放StoreBailDto
func ReleaseStoreBailDto(v *StoreBailDto) {
	v.OuterStoreId = ""
	v.IsTest = 0
	v.EtcVersion = 0
	v.Bail = 0
	v.BusinessType = 0
	v.CategoryCode = 0
	poolStoreBailDto.Put(v)
}
