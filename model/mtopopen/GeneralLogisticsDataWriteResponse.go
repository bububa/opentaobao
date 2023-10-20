package mtopopen

import (
	"sync"
)

// GeneralLogisticsDataWriteResponse 结构体
type GeneralLogisticsDataWriteResponse struct {
	// 信息写入错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 信息写入错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 信息写入是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolGeneralLogisticsDataWriteResponse = sync.Pool{
	New: func() any {
		return new(GeneralLogisticsDataWriteResponse)
	},
}

// GetGeneralLogisticsDataWriteResponse() 从对象池中获取GeneralLogisticsDataWriteResponse
func GetGeneralLogisticsDataWriteResponse() *GeneralLogisticsDataWriteResponse {
	return poolGeneralLogisticsDataWriteResponse.Get().(*GeneralLogisticsDataWriteResponse)
}

// ReleaseGeneralLogisticsDataWriteResponse 释放GeneralLogisticsDataWriteResponse
func ReleaseGeneralLogisticsDataWriteResponse(v *GeneralLogisticsDataWriteResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolGeneralLogisticsDataWriteResponse.Put(v)
}
