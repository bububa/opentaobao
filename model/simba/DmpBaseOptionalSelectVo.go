package simba

import (
	"sync"
)

// DmpBaseOptionalSelectVo 结构体
type DmpBaseOptionalSelectVo struct {
	// 数据来源app
	AppIdList []int64 `json:"app_id_list,omitempty" xml:"app_id_list>int64,omitempty"`
	// 来源展示名称
	DisplayName string `json:"display_name,omitempty" xml:"display_name,omitempty"`
}

var poolDmpBaseOptionalSelectVo = sync.Pool{
	New: func() any {
		return new(DmpBaseOptionalSelectVo)
	},
}

// GetDmpBaseOptionalSelectVo() 从对象池中获取DmpBaseOptionalSelectVo
func GetDmpBaseOptionalSelectVo() *DmpBaseOptionalSelectVo {
	return poolDmpBaseOptionalSelectVo.Get().(*DmpBaseOptionalSelectVo)
}

// ReleaseDmpBaseOptionalSelectVo 释放DmpBaseOptionalSelectVo
func ReleaseDmpBaseOptionalSelectVo(v *DmpBaseOptionalSelectVo) {
	v.AppIdList = v.AppIdList[:0]
	v.DisplayName = ""
	poolDmpBaseOptionalSelectVo.Put(v)
}
