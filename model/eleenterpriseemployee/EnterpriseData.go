package eleenterpriseemployee

import (
	"sync"
)

// EnterpriseData 结构体
type EnterpriseData struct {
	// 错误信息
	ErrorMsgs []ErrorMsg `json:"error_msgs,omitempty" xml:"error_msgs>error_msg,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolEnterpriseData = sync.Pool{
	New: func() any {
		return new(EnterpriseData)
	},
}

// GetEnterpriseData() 从对象池中获取EnterpriseData
func GetEnterpriseData() *EnterpriseData {
	return poolEnterpriseData.Get().(*EnterpriseData)
}

// ReleaseEnterpriseData 释放EnterpriseData
func ReleaseEnterpriseData(v *EnterpriseData) {
	v.ErrorMsgs = v.ErrorMsgs[:0]
	v.Success = false
	poolEnterpriseData.Put(v)
}
