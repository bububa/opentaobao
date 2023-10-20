package feedflow

import (
	"sync"
)

// LabelQueryDto 结构体
type LabelQueryDto struct {
	// 宝贝id列表
	ItemIdList []int64 `json:"item_id_list,omitempty" xml:"item_id_list>int64,omitempty"`
	// 选项值
	OptionName string `json:"option_name,omitempty" xml:"option_name,omitempty"`
	// 定向类型
	TargetType string `json:"target_type,omitempty" xml:"target_type,omitempty"`
	// 分页条件
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 定向id
	TargetId int64 `json:"target_id,omitempty" xml:"target_id,omitempty"`
	// 分页条件
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
}

var poolLabelQueryDto = sync.Pool{
	New: func() any {
		return new(LabelQueryDto)
	},
}

// GetLabelQueryDto() 从对象池中获取LabelQueryDto
func GetLabelQueryDto() *LabelQueryDto {
	return poolLabelQueryDto.Get().(*LabelQueryDto)
}

// ReleaseLabelQueryDto 释放LabelQueryDto
func ReleaseLabelQueryDto(v *LabelQueryDto) {
	v.ItemIdList = v.ItemIdList[:0]
	v.OptionName = ""
	v.TargetType = ""
	v.PageSize = 0
	v.TargetId = 0
	v.Offset = 0
	poolLabelQueryDto.Put(v)
}
