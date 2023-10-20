package user

import (
	"sync"
)

// Resultdata 结构体
type Resultdata struct {
	// 1=是AG用户，0=非AG用户
	AgAccount int64 `json:"ag_account,omitempty" xml:"ag_account,omitempty"`
}

var poolResultdata = sync.Pool{
	New: func() any {
		return new(Resultdata)
	},
}

// GetResultdata() 从对象池中获取Resultdata
func GetResultdata() *Resultdata {
	return poolResultdata.Get().(*Resultdata)
}

// ReleaseResultdata 释放Resultdata
func ReleaseResultdata(v *Resultdata) {
	v.AgAccount = 0
	poolResultdata.Put(v)
}
