package waybill

import (
	"sync"
)

// CloudPrintBaseResult 结构体
type CloudPrintBaseResult struct {
	// 数据
	Datas []CustomAreaResult `json:"datas,omitempty" xml:"datas>custom_area_result,omitempty"`
	// data
	ResourceList []IsvResourceDo `json:"resource_list,omitempty" xml:"resource_list>isv_resource_do,omitempty"`
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCloudPrintBaseResult = sync.Pool{
	New: func() any {
		return new(CloudPrintBaseResult)
	},
}

// GetCloudPrintBaseResult() 从对象池中获取CloudPrintBaseResult
func GetCloudPrintBaseResult() *CloudPrintBaseResult {
	return poolCloudPrintBaseResult.Get().(*CloudPrintBaseResult)
}

// ReleaseCloudPrintBaseResult 释放CloudPrintBaseResult
func ReleaseCloudPrintBaseResult(v *CloudPrintBaseResult) {
	v.Datas = v.Datas[:0]
	v.ResourceList = v.ResourceList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = ""
	v.Success = false
	poolCloudPrintBaseResult.Put(v)
}
