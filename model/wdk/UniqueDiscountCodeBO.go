package wdk

import (
	"sync"
)

// UniqueDiscountCodeBO 结构体
type UniqueDiscountCodeBO struct {
	// 唯一码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 过期时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
}

var poolUniqueDiscountCodeBO = sync.Pool{
	New: func() any {
		return new(UniqueDiscountCodeBO)
	},
}

// GetUniqueDiscountCodeBO() 从对象池中获取UniqueDiscountCodeBO
func GetUniqueDiscountCodeBO() *UniqueDiscountCodeBO {
	return poolUniqueDiscountCodeBO.Get().(*UniqueDiscountCodeBO)
}

// ReleaseUniqueDiscountCodeBO 释放UniqueDiscountCodeBO
func ReleaseUniqueDiscountCodeBO(v *UniqueDiscountCodeBO) {
	v.Code = ""
	v.ExpireTime = ""
	poolUniqueDiscountCodeBO.Put(v)
}
