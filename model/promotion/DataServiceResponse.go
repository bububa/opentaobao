package promotion

import (
	"sync"
)

// DataServiceResponse 结构体
type DataServiceResponse struct {
	// 报错编码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 报错描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 鹰眼id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 扩展字段
	ExtMap *Extmap `json:"ext_map,omitempty" xml:"ext_map,omitempty"`
	// 结果对象
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDataServiceResponse = sync.Pool{
	New: func() any {
		return new(DataServiceResponse)
	},
}

// GetDataServiceResponse() 从对象池中获取DataServiceResponse
func GetDataServiceResponse() *DataServiceResponse {
	return poolDataServiceResponse.Get().(*DataServiceResponse)
}

// ReleaseDataServiceResponse 释放DataServiceResponse
func ReleaseDataServiceResponse(v *DataServiceResponse) {
	v.ResultCode = ""
	v.ResultMsg = ""
	v.TraceId = ""
	v.ExtMap = nil
	v.Result = false
	v.Success = false
	poolDataServiceResponse.Put(v)
}
