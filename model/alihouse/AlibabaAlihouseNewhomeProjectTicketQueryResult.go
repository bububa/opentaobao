package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectTicketQueryResult 结构体
type AlibabaAlihouseNewhomeProjectTicketQueryResult struct {
	// 商品信息列表
	Data []ProjectVerifyTicketDto `json:"data,omitempty" xml:"data>project_verify_ticket_dto,omitempty"`
	// 返回编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 详情信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectTicketQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectTicketQueryResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectTicketQueryResult() 从对象池中获取AlibabaAlihouseNewhomeProjectTicketQueryResult
func GetAlibabaAlihouseNewhomeProjectTicketQueryResult() *AlibabaAlihouseNewhomeProjectTicketQueryResult {
	return poolAlibabaAlihouseNewhomeProjectTicketQueryResult.Get().(*AlibabaAlihouseNewhomeProjectTicketQueryResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectTicketQueryResult 释放AlibabaAlihouseNewhomeProjectTicketQueryResult
func ReleaseAlibabaAlihouseNewhomeProjectTicketQueryResult(v *AlibabaAlihouseNewhomeProjectTicketQueryResult) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectTicketQueryResult.Put(v)
}
