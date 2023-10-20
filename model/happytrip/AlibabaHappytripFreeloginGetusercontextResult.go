package happytrip

import (
	"sync"
)

// AlibabaHappytripFreeloginGetusercontextResult 结构体
type AlibabaHappytripFreeloginGetusercontextResult struct {
	// 错误消息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 返回的结果值
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
}

var poolAlibabaHappytripFreeloginGetusercontextResult = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripFreeloginGetusercontextResult)
	},
}

// GetAlibabaHappytripFreeloginGetusercontextResult() 从对象池中获取AlibabaHappytripFreeloginGetusercontextResult
func GetAlibabaHappytripFreeloginGetusercontextResult() *AlibabaHappytripFreeloginGetusercontextResult {
	return poolAlibabaHappytripFreeloginGetusercontextResult.Get().(*AlibabaHappytripFreeloginGetusercontextResult)
}

// ReleaseAlibabaHappytripFreeloginGetusercontextResult 释放AlibabaHappytripFreeloginGetusercontextResult
func ReleaseAlibabaHappytripFreeloginGetusercontextResult(v *AlibabaHappytripFreeloginGetusercontextResult) {
	v.Errmsg = ""
	v.Content = ""
	v.Errno = 0
	poolAlibabaHappytripFreeloginGetusercontextResult.Put(v)
}
