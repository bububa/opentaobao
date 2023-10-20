package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthImriskQueryAPIResponse 问诊质控接口 API返回值
// alibaba.alihealth.imrisk.query
//
// 阿里健康的问诊质控接口
type AlibabaAlihealthImriskQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthImriskQueryAPIResponseModel
}

// AlibabaAlihealthImriskQueryAPIResponseModel is 问诊质控接口 成功返回结果
type AlibabaAlihealthImriskQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_imrisk_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data []Imrule `json:"data,omitempty" xml:"data>imrule,omitempty"`
	// 错误码
	TheErrCode string `json:"the_err_code,omitempty" xml:"the_err_code,omitempty"`
	// 错误信息
	TheErrMsg string `json:"the_err_msg,omitempty" xml:"the_err_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
