package examination

import (
	"sync"
)

// ServiceResult 结构体
type ServiceResult struct {
	// 返回数据对象
	DataList []StoreSpecialTagsResponse `json:"data_list,omitempty" xml:"data_list>store_special_tags_response,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 返回数据,一般为空
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// eagleEyeTraceId
	EagleEyeTraceId string `json:"eagle_eye_trace_id,omitempty" xml:"eagle_eye_trace_id,omitempty"`
	// 执行是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolServiceResult = sync.Pool{
	New: func() any {
		return new(ServiceResult)
	},
}

// GetServiceResult() 从对象池中获取ServiceResult
func GetServiceResult() *ServiceResult {
	return poolServiceResult.Get().(*ServiceResult)
}

// ReleaseServiceResult 释放ServiceResult
func ReleaseServiceResult(v *ServiceResult) {
	v.DataList = v.DataList[:0]
	v.ErrMessage = ""
	v.Data = ""
	v.ErrCode = ""
	v.EagleEyeTraceId = ""
	v.Success = false
	poolServiceResult.Put(v)
}
