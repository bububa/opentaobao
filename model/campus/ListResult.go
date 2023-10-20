package campus

import (
	"sync"
)

// ListResult 结构体
type ListResult struct {
	// content
	Contents []string `json:"contents,omitempty" xml:"contents>string,omitempty"`
	// 设备历史数据内容
	DeviceDataList []DeviceHistoryBatchApiDto `json:"device_data_list,omitempty" xml:"device_data_list>device_history_batch_api_dto,omitempty"`
	// 规则列表数据
	ContentList []string `json:"content_list,omitempty" xml:"content_list>string,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorLevel
	ErrorLevel string `json:"error_level,omitempty" xml:"error_level,omitempty"`
	// 额外报错信息
	ErrorExtInfo string `json:"error_ext_info,omitempty" xml:"error_ext_info,omitempty"`
	// 总记录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolListResult = sync.Pool{
	New: func() any {
		return new(ListResult)
	},
}

// GetListResult() 从对象池中获取ListResult
func GetListResult() *ListResult {
	return poolListResult.Get().(*ListResult)
}

// ReleaseListResult 释放ListResult
func ReleaseListResult(v *ListResult) {
	v.Contents = v.Contents[:0]
	v.DeviceDataList = v.DeviceDataList[:0]
	v.ContentList = v.ContentList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.RequestId = ""
	v.ErrorLevel = ""
	v.ErrorExtInfo = ""
	v.TotalCount = 0
	v.Success = false
	poolListResult.Put(v)
}
