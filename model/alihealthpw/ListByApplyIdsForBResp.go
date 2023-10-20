package alihealthpw

import (
	"sync"
)

// ListByApplyIdsForBResp 结构体
type ListByApplyIdsForBResp struct {
	// 列表
	List []PendingListDto `json:"list,omitempty" xml:"list>pending_list_dto,omitempty"`
}

var poolListByApplyIdsForBResp = sync.Pool{
	New: func() any {
		return new(ListByApplyIdsForBResp)
	},
}

// GetListByApplyIdsForBResp() 从对象池中获取ListByApplyIdsForBResp
func GetListByApplyIdsForBResp() *ListByApplyIdsForBResp {
	return poolListByApplyIdsForBResp.Get().(*ListByApplyIdsForBResp)
}

// ReleaseListByApplyIdsForBResp 释放ListByApplyIdsForBResp
func ReleaseListByApplyIdsForBResp(v *ListByApplyIdsForBResp) {
	v.List = v.List[:0]
	poolListByApplyIdsForBResp.Put(v)
}
