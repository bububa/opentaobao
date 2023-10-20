package security

import (
	"sync"
)

// JaqImageDetectResultCollection 结构体
type JaqImageDetectResultCollection struct {
	// 响应消息结构体
	JaqImageDetectResultList []JaqImageDetectResult `json:"jaq_image_detect_result_list,omitempty" xml:"jaq_image_detect_result_list>jaq_image_detect_result,omitempty"`
}

var poolJaqImageDetectResultCollection = sync.Pool{
	New: func() any {
		return new(JaqImageDetectResultCollection)
	},
}

// GetJaqImageDetectResultCollection() 从对象池中获取JaqImageDetectResultCollection
func GetJaqImageDetectResultCollection() *JaqImageDetectResultCollection {
	return poolJaqImageDetectResultCollection.Get().(*JaqImageDetectResultCollection)
}

// ReleaseJaqImageDetectResultCollection 释放JaqImageDetectResultCollection
func ReleaseJaqImageDetectResultCollection(v *JaqImageDetectResultCollection) {
	v.JaqImageDetectResultList = v.JaqImageDetectResultList[:0]
	poolJaqImageDetectResultCollection.Put(v)
}
