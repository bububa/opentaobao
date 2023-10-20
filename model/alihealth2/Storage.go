package alihealth2

import (
	"sync"
)

// Storage 结构体
type Storage struct {
	// 已经使用的数量
	UsedNum string `json:"used_num,omitempty" xml:"used_num,omitempty"`
	// 日期
	StrRegdate string `json:"str_regdate,omitempty" xml:"str_regdate,omitempty"`
	// 可使用的最大数
	MaxNum string `json:"max_num,omitempty" xml:"max_num,omitempty"`
	// 状态,目前仅瑞慈使用,瑞慈状态:00:正常 01:已满 02:机构休息 03: 预约截止
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolStorage = sync.Pool{
	New: func() any {
		return new(Storage)
	},
}

// GetStorage() 从对象池中获取Storage
func GetStorage() *Storage {
	return poolStorage.Get().(*Storage)
}

// ReleaseStorage 释放Storage
func ReleaseStorage(v *Storage) {
	v.UsedNum = ""
	v.StrRegdate = ""
	v.MaxNum = ""
	v.Status = ""
	poolStorage.Put(v)
}
