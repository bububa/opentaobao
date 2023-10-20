package tmallcar

import (
	"sync"
)

// CarefreeDetailQueryReq 结构体
type CarefreeDetailQueryReq struct {
	// 业务单id
	BizId int64 `json:"biz_id,omitempty" xml:"biz_id,omitempty"`
}

var poolCarefreeDetailQueryReq = sync.Pool{
	New: func() any {
		return new(CarefreeDetailQueryReq)
	},
}

// GetCarefreeDetailQueryReq() 从对象池中获取CarefreeDetailQueryReq
func GetCarefreeDetailQueryReq() *CarefreeDetailQueryReq {
	return poolCarefreeDetailQueryReq.Get().(*CarefreeDetailQueryReq)
}

// ReleaseCarefreeDetailQueryReq 释放CarefreeDetailQueryReq
func ReleaseCarefreeDetailQueryReq(v *CarefreeDetailQueryReq) {
	v.BizId = 0
	poolCarefreeDetailQueryReq.Put(v)
}
