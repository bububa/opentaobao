package tmallgenie

import (
	"sync"
)

// AlibabaAilabsAligenieOpencontentPushResult 结构体
type AlibabaAilabsAligenieOpencontentPushResult struct {
	// retMsg
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// retCode
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

var poolAlibabaAilabsAligenieOpencontentPushResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieOpencontentPushResult)
	},
}

// GetAlibabaAilabsAligenieOpencontentPushResult() 从对象池中获取AlibabaAilabsAligenieOpencontentPushResult
func GetAlibabaAilabsAligenieOpencontentPushResult() *AlibabaAilabsAligenieOpencontentPushResult {
	return poolAlibabaAilabsAligenieOpencontentPushResult.Get().(*AlibabaAilabsAligenieOpencontentPushResult)
}

// ReleaseAlibabaAilabsAligenieOpencontentPushResult 释放AlibabaAilabsAligenieOpencontentPushResult
func ReleaseAlibabaAilabsAligenieOpencontentPushResult(v *AlibabaAilabsAligenieOpencontentPushResult) {
	v.RetMsg = ""
	v.RetCode = 0
	poolAlibabaAilabsAligenieOpencontentPushResult.Put(v)
}
