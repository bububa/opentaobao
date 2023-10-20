package scbp

import (
	"sync"
)

// BidStrategyDto 结构体
type BidStrategyDto struct {
	// 主键
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 目标排名，前N名
	Topn int64 `json:"topn,omitempty" xml:"topn,omitempty"`
	// 出价间隔分钟为单位的数值型
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 产品
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 溢价比例
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
}

var poolBidStrategyDto = sync.Pool{
	New: func() any {
		return new(BidStrategyDto)
	},
}

// GetBidStrategyDto() 从对象池中获取BidStrategyDto
func GetBidStrategyDto() *BidStrategyDto {
	return poolBidStrategyDto.Get().(*BidStrategyDto)
}

// ReleaseBidStrategyDto 释放BidStrategyDto
func ReleaseBidStrategyDto(v *BidStrategyDto) {
	v.Id = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Topn = 0
	v.Duration = 0
	v.ProductId = 0
	v.Discount = 0
	poolBidStrategyDto.Put(v)
}
