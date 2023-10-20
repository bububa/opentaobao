package security

import (
	"sync"
)

// JaqOcrImageDetectResult 结构体
type JaqOcrImageDetectResult struct {
	// 字符串列表，内容是图像中文字的主要段落内容（按照概率输出最多5组概率最大的组合）
	Texts []string `json:"texts,omitempty" xml:"texts>string,omitempty"`
}

var poolJaqOcrImageDetectResult = sync.Pool{
	New: func() any {
		return new(JaqOcrImageDetectResult)
	},
}

// GetJaqOcrImageDetectResult() 从对象池中获取JaqOcrImageDetectResult
func GetJaqOcrImageDetectResult() *JaqOcrImageDetectResult {
	return poolJaqOcrImageDetectResult.Get().(*JaqOcrImageDetectResult)
}

// ReleaseJaqOcrImageDetectResult 释放JaqOcrImageDetectResult
func ReleaseJaqOcrImageDetectResult(v *JaqOcrImageDetectResult) {
	v.Texts = v.Texts[:0]
	poolJaqOcrImageDetectResult.Put(v)
}
