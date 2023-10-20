package alihouse

import (
	"sync"
)

// AstoreSceneInfoDto 结构体
type AstoreSceneInfoDto struct {
	// 外部主体id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 外部商家id
	OuterSellerId string `json:"outer_seller_id,omitempty" xml:"outer_seller_id,omitempty"`
	// 外部主体类型
	OuterType int64 `json:"outer_type,omitempty" xml:"outer_type,omitempty"`
}

var poolAstoreSceneInfoDto = sync.Pool{
	New: func() any {
		return new(AstoreSceneInfoDto)
	},
}

// GetAstoreSceneInfoDto() 从对象池中获取AstoreSceneInfoDto
func GetAstoreSceneInfoDto() *AstoreSceneInfoDto {
	return poolAstoreSceneInfoDto.Get().(*AstoreSceneInfoDto)
}

// ReleaseAstoreSceneInfoDto 释放AstoreSceneInfoDto
func ReleaseAstoreSceneInfoDto(v *AstoreSceneInfoDto) {
	v.OuterId = ""
	v.OuterStoreId = ""
	v.OuterSellerId = ""
	v.OuterType = 0
	poolAstoreSceneInfoDto.Put(v)
}
