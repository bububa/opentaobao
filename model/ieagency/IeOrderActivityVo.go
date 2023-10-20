package ieagency

import (
	"sync"
)

// IeOrderActivityVo 结构体
type IeOrderActivityVo struct {
	// activityName
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// activityPrice
	ActivityPrice int64 `json:"activity_price,omitempty" xml:"activity_price,omitempty"`
}

var poolIeOrderActivityVo = sync.Pool{
	New: func() any {
		return new(IeOrderActivityVo)
	},
}

// GetIeOrderActivityVo() 从对象池中获取IeOrderActivityVo
func GetIeOrderActivityVo() *IeOrderActivityVo {
	return poolIeOrderActivityVo.Get().(*IeOrderActivityVo)
}

// ReleaseIeOrderActivityVo 释放IeOrderActivityVo
func ReleaseIeOrderActivityVo(v *IeOrderActivityVo) {
	v.ActivityName = ""
	v.ActivityPrice = 0
	poolIeOrderActivityVo.Put(v)
}
