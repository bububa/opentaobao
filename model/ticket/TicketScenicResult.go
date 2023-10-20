package ticket

import (
	"sync"
)

// TicketScenicResult 结构体
type TicketScenicResult struct {
	// 商家景点编码
	OutScenicId string `json:"out_scenic_id,omitempty" xml:"out_scenic_id,omitempty"`
	// 扩展字段，预留
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 阿里标准景点库ID
	AliScenicId int64 `json:"ali_scenic_id,omitempty" xml:"ali_scenic_id,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTicketScenicResult = sync.Pool{
	New: func() any {
		return new(TicketScenicResult)
	},
}

// GetTicketScenicResult() 从对象池中获取TicketScenicResult
func GetTicketScenicResult() *TicketScenicResult {
	return poolTicketScenicResult.Get().(*TicketScenicResult)
}

// ReleaseTicketScenicResult 释放TicketScenicResult
func ReleaseTicketScenicResult(v *TicketScenicResult) {
	v.OutScenicId = ""
	v.Extend = ""
	v.AliScenicId = 0
	v.Success = false
	poolTicketScenicResult.Put(v)
}
