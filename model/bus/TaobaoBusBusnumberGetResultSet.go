package bus

import (
	"sync"
)

// TaobaoBusBusnumberGetResultSet 结构体
type TaobaoBusBusnumberGetResultSet struct {
	// errCode BUSNUMBER_SEARCH_NOBUS 找不到班次		POWER_D权限问题
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// module
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoBusBusnumberGetResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoBusBusnumberGetResultSet)
	},
}

// GetTaobaoBusBusnumberGetResultSet() 从对象池中获取TaobaoBusBusnumberGetResultSet
func GetTaobaoBusBusnumberGetResultSet() *TaobaoBusBusnumberGetResultSet {
	return poolTaobaoBusBusnumberGetResultSet.Get().(*TaobaoBusBusnumberGetResultSet)
}

// ReleaseTaobaoBusBusnumberGetResultSet 释放TaobaoBusBusnumberGetResultSet
func ReleaseTaobaoBusBusnumberGetResultSet(v *TaobaoBusBusnumberGetResultSet) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Module = ""
	v.Success = false
	poolTaobaoBusBusnumberGetResultSet.Put(v)
}
