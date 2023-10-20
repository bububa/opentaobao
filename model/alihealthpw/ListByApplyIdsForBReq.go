package alihealthpw

import (
	"sync"
)

// ListByApplyIdsForBReq 结构体
type ListByApplyIdsForBReq struct {
	// 编码列表
	ApplyUniqueCodes []string `json:"apply_unique_codes,omitempty" xml:"apply_unique_codes>string,omitempty"`
}

var poolListByApplyIdsForBReq = sync.Pool{
	New: func() any {
		return new(ListByApplyIdsForBReq)
	},
}

// GetListByApplyIdsForBReq() 从对象池中获取ListByApplyIdsForBReq
func GetListByApplyIdsForBReq() *ListByApplyIdsForBReq {
	return poolListByApplyIdsForBReq.Get().(*ListByApplyIdsForBReq)
}

// ReleaseListByApplyIdsForBReq 释放ListByApplyIdsForBReq
func ReleaseListByApplyIdsForBReq(v *ListByApplyIdsForBReq) {
	v.ApplyUniqueCodes = v.ApplyUniqueCodes[:0]
	poolListByApplyIdsForBReq.Put(v)
}
