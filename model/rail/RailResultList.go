package rail

import (
	"sync"
)

// RailResultList 结构体
type RailResultList struct {
	// 城市列表
	ModuleList []RailDivisionRs `json:"module_list,omitempty" xml:"module_list>rail_division_rs,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 错误描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
}

var poolRailResultList = sync.Pool{
	New: func() any {
		return new(RailResultList)
	},
}

// GetRailResultList() 从对象池中获取RailResultList
func GetRailResultList() *RailResultList {
	return poolRailResultList.Get().(*RailResultList)
}

// ReleaseRailResultList 释放RailResultList
func ReleaseRailResultList(v *RailResultList) {
	v.ModuleList = v.ModuleList[:0]
	v.Success = ""
	v.ErrMsg = ""
	v.ErrCode = ""
	poolRailResultList.Put(v)
}
