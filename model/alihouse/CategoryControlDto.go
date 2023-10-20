package alihouse

import (
	"sync"
)

// CategoryControlDto 结构体
type CategoryControlDto struct {
	// 外部id列表
	OuterTargetIds []string `json:"outer_target_ids,omitempty" xml:"outer_target_ids>string,omitempty"`
	// 二级类目列表
	SecondCategory []int64 `json:"second_category,omitempty" xml:"second_category>int64,omitempty"`
	// 三级类目列表
	ThirdCategory []int64 `json:"third_category,omitempty" xml:"third_category>int64,omitempty"`
	// 外部id
	OuterTargetId string `json:"outer_target_id,omitempty" xml:"outer_target_id,omitempty"`
	// 1
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 是否测试
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 类型
	TargetType int64 `json:"target_type,omitempty" xml:"target_type,omitempty"`
	// 操作类型 0-删除 1-新增
	OperateType int64 `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
}

var poolCategoryControlDto = sync.Pool{
	New: func() any {
		return new(CategoryControlDto)
	},
}

// GetCategoryControlDto() 从对象池中获取CategoryControlDto
func GetCategoryControlDto() *CategoryControlDto {
	return poolCategoryControlDto.Get().(*CategoryControlDto)
}

// ReleaseCategoryControlDto 释放CategoryControlDto
func ReleaseCategoryControlDto(v *CategoryControlDto) {
	v.OuterTargetIds = v.OuterTargetIds[:0]
	v.SecondCategory = v.SecondCategory[:0]
	v.ThirdCategory = v.ThirdCategory[:0]
	v.OuterTargetId = ""
	v.OuterStoreId = ""
	v.IsTest = 0
	v.TargetType = 0
	v.OperateType = 0
	poolCategoryControlDto.Put(v)
}
