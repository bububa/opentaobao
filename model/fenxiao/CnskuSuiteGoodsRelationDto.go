package fenxiao

import (
	"sync"
)

// CnskuSuiteGoodsRelationDto 结构体
type CnskuSuiteGoodsRelationDto struct {
	// 真实货值
	GoodsValue string `json:"goods_value,omitempty" xml:"goods_value,omitempty"`
	// 主品Id，创建可不填
	BelongSuiteGoodId int64 `json:"belong_suite_good_id,omitempty" xml:"belong_suite_good_id,omitempty"`
	// 单子品数量，要求&gt;0
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 成分子品Id
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
}

var poolCnskuSuiteGoodsRelationDto = sync.Pool{
	New: func() any {
		return new(CnskuSuiteGoodsRelationDto)
	},
}

// GetCnskuSuiteGoodsRelationDto() 从对象池中获取CnskuSuiteGoodsRelationDto
func GetCnskuSuiteGoodsRelationDto() *CnskuSuiteGoodsRelationDto {
	return poolCnskuSuiteGoodsRelationDto.Get().(*CnskuSuiteGoodsRelationDto)
}

// ReleaseCnskuSuiteGoodsRelationDto 释放CnskuSuiteGoodsRelationDto
func ReleaseCnskuSuiteGoodsRelationDto(v *CnskuSuiteGoodsRelationDto) {
	v.GoodsValue = ""
	v.BelongSuiteGoodId = 0
	v.Amount = 0
	v.GoodsId = 0
	poolCnskuSuiteGoodsRelationDto.Put(v)
}
