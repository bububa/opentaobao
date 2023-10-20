package tmallsc

import (
	"sync"
)

// TmallWorkcardIdentifyResult 结构体
type TmallWorkcardIdentifyResult struct {
	// gmtModified
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 返回void
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallWorkcardIdentifyResult = sync.Pool{
	New: func() any {
		return new(TmallWorkcardIdentifyResult)
	},
}

// GetTmallWorkcardIdentifyResult() 从对象池中获取TmallWorkcardIdentifyResult
func GetTmallWorkcardIdentifyResult() *TmallWorkcardIdentifyResult {
	return poolTmallWorkcardIdentifyResult.Get().(*TmallWorkcardIdentifyResult)
}

// ReleaseTmallWorkcardIdentifyResult 释放TmallWorkcardIdentifyResult
func ReleaseTmallWorkcardIdentifyResult(v *TmallWorkcardIdentifyResult) {
	v.GmtModified = ""
	v.GmtCreate = ""
	v.Value = ""
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Success = false
	poolTmallWorkcardIdentifyResult.Put(v)
}
