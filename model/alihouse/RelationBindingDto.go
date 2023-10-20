package alihouse

import (
	"sync"
)

// RelationBindingDto 结构体
type RelationBindingDto struct {
	// 货下挂的其他品列表 最大列表长度：100
	Extend []GoodsRelationDto `json:"extend,omitempty" xml:"extend>goods_relation_dto,omitempty"`
	// 货下挂的其他货列表 最大列表长度：100
	RelationCargos []GoodsRelationDto `json:"relation_cargos,omitempty" xml:"relation_cargos>goods_relation_dto,omitempty"`
	// 外部货id（外部唯一识别码）
	OuterTid string `json:"outer_tid,omitempty" xml:"outer_tid,omitempty"`
	// 外部私域楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 货所属外部项目店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
}

var poolRelationBindingDto = sync.Pool{
	New: func() any {
		return new(RelationBindingDto)
	},
}

// GetRelationBindingDto() 从对象池中获取RelationBindingDto
func GetRelationBindingDto() *RelationBindingDto {
	return poolRelationBindingDto.Get().(*RelationBindingDto)
}

// ReleaseRelationBindingDto 释放RelationBindingDto
func ReleaseRelationBindingDto(v *RelationBindingDto) {
	v.Extend = v.Extend[:0]
	v.RelationCargos = v.RelationCargos[:0]
	v.OuterTid = ""
	v.OuterId = ""
	v.OuterStoreId = ""
	poolRelationBindingDto.Put(v)
}
