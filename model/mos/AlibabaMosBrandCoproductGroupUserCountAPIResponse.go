package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosbrandcoproductgroupusercountAPIResponse 按照查询条件统计总数 API返回值
// alibaba.mos.brand.coproduct.group.user.count
//
// 按照查询条件统计总数
type AlibabamosbrandcoproductgroupusercountAPIResponse struct {
	model.CommonResponse
	AlibabamosbrandcoproductgroupusercountAPIResponseModel
}

// AlibabamosbrandcoproductgroupusercountAPIResponseModel is 按照查询条件统计总数 成功返回结果
type AlibabamosbrandcoproductgroupusercountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_brand_coproduct_group_user_count_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用链id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 异步结果
	AsyncResult string `json:"async_result,omitempty" xml:"async_result,omitempty"`
	// 错误code
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 扩展字段
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 统计数量
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Issuccess bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
	// 是否同步
	IsAsync bool `json:"is_async,omitempty" xml:"is_async,omitempty"`
}
