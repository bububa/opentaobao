package alihouse

import (
	"sync"
)

// AstoreShopConfigDto 结构体
type AstoreShopConfigDto struct {
	// 外部ID
	OuterTargetId string `json:"outer_target_id,omitempty" xml:"outer_target_id,omitempty"`
	// 外部门店ID（为了找小B使用）
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 外部类型
	TargetType int64 `json:"target_type,omitempty" xml:"target_type,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 装修类型
	FitupType int64 `json:"fitup_type,omitempty" xml:"fitup_type,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolAstoreShopConfigDto = sync.Pool{
	New: func() any {
		return new(AstoreShopConfigDto)
	},
}

// GetAstoreShopConfigDto() 从对象池中获取AstoreShopConfigDto
func GetAstoreShopConfigDto() *AstoreShopConfigDto {
	return poolAstoreShopConfigDto.Get().(*AstoreShopConfigDto)
}

// ReleaseAstoreShopConfigDto 释放AstoreShopConfigDto
func ReleaseAstoreShopConfigDto(v *AstoreShopConfigDto) {
	v.OuterTargetId = ""
	v.OuterStoreId = ""
	v.TargetType = 0
	v.Status = 0
	v.FitupType = 0
	v.Version = 0
	poolAstoreShopConfigDto.Put(v)
}
