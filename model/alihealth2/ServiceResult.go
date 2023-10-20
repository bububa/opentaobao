package alihealth2

import (
	"sync"
)

// ServiceResult 结构体
type ServiceResult struct {
	// 结果集
	Datas []BaseRule `json:"datas,omitempty" xml:"datas>base_rule,omitempty"`
	// data
	DivisionList []DivisionDto `json:"division_list,omitempty" xml:"division_list>division_dto,omitempty"`
	// msg_info
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msg_code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// errMessage
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回数据对象
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// success
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
	v.Datas = v.Datas[:0]
	v.DivisionList = v.DivisionList[:0]
	v.MsgInfo = ""
	v.MsgCode = ""
	v.ErrMessage = ""
	v.ErrCode = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = false
	v.Success = false
	poolServiceResult.Put(v)
}
