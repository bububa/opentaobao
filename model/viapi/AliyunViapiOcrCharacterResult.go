package viapi

import (
	"sync"
)

// AliyunViapiOcrCharacterResult 结构体
type AliyunViapiOcrCharacterResult struct {
	// 文字内容
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 文字区域
	TextRectangle *TextRectangle `json:"text_rectangle,omitempty" xml:"text_rectangle,omitempty"`
	// 文字区域概率，概率值的范围为[0, 1]
	Probability int64 `json:"probability,omitempty" xml:"probability,omitempty"`
}

var poolAliyunViapiOcrCharacterResult = sync.Pool{
	New: func() any {
		return new(AliyunViapiOcrCharacterResult)
	},
}

// GetAliyunViapiOcrCharacterResult() 从对象池中获取AliyunViapiOcrCharacterResult
func GetAliyunViapiOcrCharacterResult() *AliyunViapiOcrCharacterResult {
	return poolAliyunViapiOcrCharacterResult.Get().(*AliyunViapiOcrCharacterResult)
}

// ReleaseAliyunViapiOcrCharacterResult 释放AliyunViapiOcrCharacterResult
func ReleaseAliyunViapiOcrCharacterResult(v *AliyunViapiOcrCharacterResult) {
	v.Text = ""
	v.TextRectangle = nil
	v.Probability = 0
	poolAliyunViapiOcrCharacterResult.Put(v)
}
