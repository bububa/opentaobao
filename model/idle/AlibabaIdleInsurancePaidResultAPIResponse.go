package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleInsurancePaidResultAPIResponse 闲鱼行业保险理赔结果同步 API返回值
// alibaba.idle.insurance.paid.result
//
// 闲鱼行业保险理赔结果同步
type AlibabaIdleInsurancePaidResultAPIResponse struct {
	model.CommonResponse
	AlibabaIdleInsurancePaidResultAPIResponseModel
}

// AlibabaIdleInsurancePaidResultAPIResponseModel is 闲鱼行业保险理赔结果同步 成功返回结果
type AlibabaIdleInsurancePaidResultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_insurance_paid_result_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误code
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 错误msg
	EMsg string `json:"e_msg,omitempty" xml:"e_msg,omitempty"`
	// 同步是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
