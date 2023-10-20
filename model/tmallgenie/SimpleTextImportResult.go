package tmallgenie

import (
	"sync"
)

// SimpleTextImportResult 结构体
type SimpleTextImportResult struct {
	// 已存在或重复的实体数
	NumExist int64 `json:"num_exist,omitempty" xml:"num_exist,omitempty"`
	// 上传失败的实体数
	NumFailed int64 `json:"num_failed,omitempty" xml:"num_failed,omitempty"`
	// 上传成功的实体数
	NumSuccessful int64 `json:"num_successful,omitempty" xml:"num_successful,omitempty"`
}

var poolSimpleTextImportResult = sync.Pool{
	New: func() any {
		return new(SimpleTextImportResult)
	},
}

// GetSimpleTextImportResult() 从对象池中获取SimpleTextImportResult
func GetSimpleTextImportResult() *SimpleTextImportResult {
	return poolSimpleTextImportResult.Get().(*SimpleTextImportResult)
}

// ReleaseSimpleTextImportResult 释放SimpleTextImportResult
func ReleaseSimpleTextImportResult(v *SimpleTextImportResult) {
	v.NumExist = 0
	v.NumFailed = 0
	v.NumSuccessful = 0
	poolSimpleTextImportResult.Put(v)
}
