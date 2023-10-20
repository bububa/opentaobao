package viapi

import (
	"sync"
)

// Element 结构体
type Element struct {
	// 单个文本的检测结果
	Results []AliyunViapiImageauditScantextResult `json:"results,omitempty" xml:"results>aliyun_viapi_imageaudit_scantext_result,omitempty"`
	// 任务Id
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

var poolElement = sync.Pool{
	New: func() any {
		return new(Element)
	},
}

// GetElement() 从对象池中获取Element
func GetElement() *Element {
	return poolElement.Get().(*Element)
}

// ReleaseElement 释放Element
func ReleaseElement(v *Element) {
	v.Results = v.Results[:0]
	v.TaskId = ""
	poolElement.Put(v)
}
