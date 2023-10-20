package charity

import (
	"sync"
)

// BillDto 结构体
type BillDto struct {
	// 账单编号
	BillId string `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
	// 账期
	BillTime int64 `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
}

var poolBillDto = sync.Pool{
	New: func() any {
		return new(BillDto)
	},
}

// GetBillDto() 从对象池中获取BillDto
func GetBillDto() *BillDto {
	return poolBillDto.Get().(*BillDto)
}

// ReleaseBillDto 释放BillDto
func ReleaseBillDto(v *BillDto) {
	v.BillId = ""
	v.BillTime = 0
	poolBillDto.Put(v)
}
