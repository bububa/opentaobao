package idle

import (
	"sync"
)

// IdleAdvResult 结构体
type IdleAdvResult struct {
	// 部分失败的素材错误原因细节列表，如果素材上传成功，不会在该列表中
	ErrDetail []IdleAdvMaterialUploadTopResult `json:"err_detail,omitempty" xml:"err_detail>idle_adv_material_upload_top_result,omitempty"`
	// 接口层面错误的错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 接口层面错误的错误原因描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口层面是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolIdleAdvResult = sync.Pool{
	New: func() any {
		return new(IdleAdvResult)
	},
}

// GetIdleAdvResult() 从对象池中获取IdleAdvResult
func GetIdleAdvResult() *IdleAdvResult {
	return poolIdleAdvResult.Get().(*IdleAdvResult)
}

// ReleaseIdleAdvResult 释放IdleAdvResult
func ReleaseIdleAdvResult(v *IdleAdvResult) {
	v.ErrDetail = v.ErrDetail[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolIdleAdvResult.Put(v)
}
