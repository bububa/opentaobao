package tmallnr

import (
	"sync"
)

// AlibabaLsyCrmActivityPageUpdateResultDo 结构体
type AlibabaLsyCrmActivityPageUpdateResultDo struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 活动ID
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLsyCrmActivityPageUpdateResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmActivityPageUpdateResultDo)
	},
}

// GetAlibabaLsyCrmActivityPageUpdateResultDo() 从对象池中获取AlibabaLsyCrmActivityPageUpdateResultDo
func GetAlibabaLsyCrmActivityPageUpdateResultDo() *AlibabaLsyCrmActivityPageUpdateResultDo {
	return poolAlibabaLsyCrmActivityPageUpdateResultDo.Get().(*AlibabaLsyCrmActivityPageUpdateResultDo)
}

// ReleaseAlibabaLsyCrmActivityPageUpdateResultDo 释放AlibabaLsyCrmActivityPageUpdateResultDo
func ReleaseAlibabaLsyCrmActivityPageUpdateResultDo(v *AlibabaLsyCrmActivityPageUpdateResultDo) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Data = 0
	v.Success = false
	poolAlibabaLsyCrmActivityPageUpdateResultDo.Put(v)
}
