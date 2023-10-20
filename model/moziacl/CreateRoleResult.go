package moziacl

import (
	"sync"
)

// CreateRoleResult 结构体
type CreateRoleResult struct {
	// 创建角色返回data，此处无数据返回
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 请求唯一id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应message，若失败则返回失败原因
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 扩展字段，与入参扩展字段值相同
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 响应code
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 接口调用是否成功，若成功则为true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCreateRoleResult = sync.Pool{
	New: func() any {
		return new(CreateRoleResult)
	},
}

// GetCreateRoleResult() 从对象池中获取CreateRoleResult
func GetCreateRoleResult() *CreateRoleResult {
	return poolCreateRoleResult.Get().(*CreateRoleResult)
}

// ReleaseCreateRoleResult 释放CreateRoleResult
func ReleaseCreateRoleResult(v *CreateRoleResult) {
	v.Data = ""
	v.RequestId = ""
	v.ResponseMessage = ""
	v.ResponseMetaData = ""
	v.ResponseCode = ""
	v.Success = false
	poolCreateRoleResult.Put(v)
}
