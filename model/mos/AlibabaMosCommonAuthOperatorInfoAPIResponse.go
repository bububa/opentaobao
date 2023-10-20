package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoscommonauthoperatorinfoAPIResponse 获取当前人员信息 API返回值
// alibaba.mos.common.auth.operator.info
//
// 获取当前人员信息
type AlibabamoscommonauthoperatorinfoAPIResponse struct {
	model.CommonResponse
	AlibabamoscommonauthoperatorinfoAPIResponseModel
}

// AlibabamoscommonauthoperatorinfoAPIResponseModel is 获取当前人员信息 成功返回结果
type AlibabamoscommonauthoperatorinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_common_auth_operator_info_response"`
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
	// 授权的人员信息
	Data *OperatorUserInfo `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Issuccess bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
	// 是否同步
	IsAsync bool `json:"is_async,omitempty" xml:"is_async,omitempty"`
}
