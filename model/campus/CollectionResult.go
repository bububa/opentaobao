package campus

import (
	"sync"
)

// CollectionResult 结构体
type CollectionResult struct {
	// 菜单内容
	Contents []TreeNode `json:"contents,omitempty" xml:"contents>tree_node,omitempty"`
	// content
	ContentList []AlibabaCampusCoreAppGetappusagesT `json:"content_list,omitempty" xml:"content_list>alibaba_campus_core_app_getappusages_t,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误级别
	ErrorLevel string `json:"error_level,omitempty" xml:"error_level,omitempty"`
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorExtInfo
	ErrorExtInfo string `json:"error_ext_info,omitempty" xml:"error_ext_info,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCollectionResult = sync.Pool{
	New: func() any {
		return new(CollectionResult)
	},
}

// GetCollectionResult() 从对象池中获取CollectionResult
func GetCollectionResult() *CollectionResult {
	return poolCollectionResult.Get().(*CollectionResult)
}

// ReleaseCollectionResult 释放CollectionResult
func ReleaseCollectionResult(v *CollectionResult) {
	v.Contents = v.Contents[:0]
	v.ContentList = v.ContentList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.ErrorLevel = ""
	v.RequestId = ""
	v.ErrorExtInfo = ""
	v.TotalCount = 0
	v.Success = false
	poolCollectionResult.Put(v)
}
