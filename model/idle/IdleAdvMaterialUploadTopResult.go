package idle

import (
	"sync"
)

// IdleAdvMaterialUploadTopResult 结构体
type IdleAdvMaterialUploadTopResult struct {
	// 失败的素材的名称，和参数传入的title对应
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 当前错误原因
	Result *IdleAdvBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

var poolIdleAdvMaterialUploadTopResult = sync.Pool{
	New: func() any {
		return new(IdleAdvMaterialUploadTopResult)
	},
}

// GetIdleAdvMaterialUploadTopResult() 从对象池中获取IdleAdvMaterialUploadTopResult
func GetIdleAdvMaterialUploadTopResult() *IdleAdvMaterialUploadTopResult {
	return poolIdleAdvMaterialUploadTopResult.Get().(*IdleAdvMaterialUploadTopResult)
}

// ReleaseIdleAdvMaterialUploadTopResult 释放IdleAdvMaterialUploadTopResult
func ReleaseIdleAdvMaterialUploadTopResult(v *IdleAdvMaterialUploadTopResult) {
	v.Title = ""
	v.Result = nil
	poolIdleAdvMaterialUploadTopResult.Put(v)
}
