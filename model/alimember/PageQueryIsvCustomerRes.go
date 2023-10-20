package alimember

import (
	"sync"
)

// PageQueryIsvCustomerRes 结构体
type PageQueryIsvCustomerRes struct {
	// 会员数据
	Data []SaasCustomerInfo `json:"data,omitempty" xml:"data>saas_customer_info,omitempty"`
	// 总大小
	TotalSize int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
	// 当前第几页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否还有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

var poolPageQueryIsvCustomerRes = sync.Pool{
	New: func() any {
		return new(PageQueryIsvCustomerRes)
	},
}

// GetPageQueryIsvCustomerRes() 从对象池中获取PageQueryIsvCustomerRes
func GetPageQueryIsvCustomerRes() *PageQueryIsvCustomerRes {
	return poolPageQueryIsvCustomerRes.Get().(*PageQueryIsvCustomerRes)
}

// ReleasePageQueryIsvCustomerRes 释放PageQueryIsvCustomerRes
func ReleasePageQueryIsvCustomerRes(v *PageQueryIsvCustomerRes) {
	v.Data = v.Data[:0]
	v.TotalSize = 0
	v.PageNo = 0
	v.PageSize = 0
	v.HasNext = false
	poolPageQueryIsvCustomerRes.Put(v)
}
