package drugtrace

import (
	"sync"
)

// ResProdCodes 结构体
type ResProdCodes struct {
	// 资源码
	ResCodeList []ResCode `json:"res_code_list,omitempty" xml:"res_code_list>res_code,omitempty"`
}

var poolResProdCodes = sync.Pool{
	New: func() any {
		return new(ResProdCodes)
	},
}

// GetResProdCodes() 从对象池中获取ResProdCodes
func GetResProdCodes() *ResProdCodes {
	return poolResProdCodes.Get().(*ResProdCodes)
}

// ReleaseResProdCodes 释放ResProdCodes
func ReleaseResProdCodes(v *ResProdCodes) {
	v.ResCodeList = v.ResCodeList[:0]
	poolResProdCodes.Put(v)
}
