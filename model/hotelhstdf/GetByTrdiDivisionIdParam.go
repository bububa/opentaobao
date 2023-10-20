package hotelhstdf

import (
	"sync"
)

// GetByTrdiDivisionIdParam 结构体
type GetByTrdiDivisionIdParam struct {
	// 第1页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页取100条数据
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 平台行政区划id，北京市
	TrdiDivisionId int64 `json:"trdi_division_id,omitempty" xml:"trdi_division_id,omitempty"`
}

var poolGetByTrdiDivisionIdParam = sync.Pool{
	New: func() any {
		return new(GetByTrdiDivisionIdParam)
	},
}

// GetGetByTrdiDivisionIdParam() 从对象池中获取GetByTrdiDivisionIdParam
func GetGetByTrdiDivisionIdParam() *GetByTrdiDivisionIdParam {
	return poolGetByTrdiDivisionIdParam.Get().(*GetByTrdiDivisionIdParam)
}

// ReleaseGetByTrdiDivisionIdParam 释放GetByTrdiDivisionIdParam
func ReleaseGetByTrdiDivisionIdParam(v *GetByTrdiDivisionIdParam) {
	v.Page = 0
	v.PageSize = 0
	v.TrdiDivisionId = 0
	poolGetByTrdiDivisionIdParam.Put(v)
}
