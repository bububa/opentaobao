package product

import (
	"sync"
)

// DisplayQualifications 结构体
type DisplayQualifications struct {
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 填充列表
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolDisplayQualifications = sync.Pool{
	New: func() any {
		return new(DisplayQualifications)
	},
}

// GetDisplayQualifications() 从对象池中获取DisplayQualifications
func GetDisplayQualifications() *DisplayQualifications {
	return poolDisplayQualifications.Get().(*DisplayQualifications)
}

// ReleaseDisplayQualifications 释放DisplayQualifications
func ReleaseDisplayQualifications(v *DisplayQualifications) {
	v.Message = ""
	v.Model = ""
	v.Result = false
	poolDisplayQualifications.Put(v)
}
