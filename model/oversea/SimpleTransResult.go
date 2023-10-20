package oversea

import (
	"sync"
)

// SimpleTransResult 结构体
type SimpleTransResult struct {
	// translatedText
	TranslatedText string `json:"translated_text,omitempty" xml:"translated_text,omitempty"`
	// statusCode
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

var poolSimpleTransResult = sync.Pool{
	New: func() any {
		return new(SimpleTransResult)
	},
}

// GetSimpleTransResult() 从对象池中获取SimpleTransResult
func GetSimpleTransResult() *SimpleTransResult {
	return poolSimpleTransResult.Get().(*SimpleTransResult)
}

// ReleaseSimpleTransResult 释放SimpleTransResult
func ReleaseSimpleTransResult(v *SimpleTransResult) {
	v.TranslatedText = ""
	v.StatusCode = ""
	poolSimpleTransResult.Put(v)
}
