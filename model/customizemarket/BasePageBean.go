package customizemarket

import (
	"sync"
)

// BasePageBean 结构体
type BasePageBean struct {
	// 数据
	Data []PetProfileDto `json:"data,omitempty" xml:"data>pet_profile_dto,omitempty"`
	// 当前页面
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总行数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolBasePageBean = sync.Pool{
	New: func() any {
		return new(BasePageBean)
	},
}

// GetBasePageBean() 从对象池中获取BasePageBean
func GetBasePageBean() *BasePageBean {
	return poolBasePageBean.Get().(*BasePageBean)
}

// ReleaseBasePageBean 释放BasePageBean
func ReleaseBasePageBean(v *BasePageBean) {
	v.Data = v.Data[:0]
	v.CurrentPage = 0
	v.TotalCount = 0
	poolBasePageBean.Put(v)
}
