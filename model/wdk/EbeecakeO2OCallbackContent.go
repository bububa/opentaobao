package wdk

import (
	"sync"
)

// EbeecakeO2OCallbackContent 结构体
type EbeecakeO2OCallbackContent struct {
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 作业内容单号
	WorkUnitContentId string `json:"work_unit_content_id,omitempty" xml:"work_unit_content_id,omitempty"`
}

var poolEbeecakeO2OCallbackContent = sync.Pool{
	New: func() any {
		return new(EbeecakeO2OCallbackContent)
	},
}

// GetEbeecakeO2OCallbackContent() 从对象池中获取EbeecakeO2OCallbackContent
func GetEbeecakeO2OCallbackContent() *EbeecakeO2OCallbackContent {
	return poolEbeecakeO2OCallbackContent.Get().(*EbeecakeO2OCallbackContent)
}

// ReleaseEbeecakeO2OCallbackContent 释放EbeecakeO2OCallbackContent
func ReleaseEbeecakeO2OCallbackContent(v *EbeecakeO2OCallbackContent) {
	v.Status = ""
	v.WorkUnitContentId = ""
	poolEbeecakeO2OCallbackContent.Put(v)
}
