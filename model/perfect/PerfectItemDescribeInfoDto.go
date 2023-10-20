package perfect

import (
	"sync"
)

// PerfectItemDescribeInfoDto 结构体
type PerfectItemDescribeInfoDto struct {
	// 商品图片
	ItemPictures []string `json:"item_pictures,omitempty" xml:"item_pictures>string,omitempty"`
	// 详情描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
}

var poolPerfectItemDescribeInfoDto = sync.Pool{
	New: func() any {
		return new(PerfectItemDescribeInfoDto)
	},
}

// GetPerfectItemDescribeInfoDto() 从对象池中获取PerfectItemDescribeInfoDto
func GetPerfectItemDescribeInfoDto() *PerfectItemDescribeInfoDto {
	return poolPerfectItemDescribeInfoDto.Get().(*PerfectItemDescribeInfoDto)
}

// ReleasePerfectItemDescribeInfoDto 释放PerfectItemDescribeInfoDto
func ReleasePerfectItemDescribeInfoDto(v *PerfectItemDescribeInfoDto) {
	v.ItemPictures = v.ItemPictures[:0]
	v.Description = ""
	v.Title = ""
	poolPerfectItemDescribeInfoDto.Put(v)
}
