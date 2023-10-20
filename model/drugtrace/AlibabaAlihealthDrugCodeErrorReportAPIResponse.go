package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeErrorReportAPIResponse 码信息错误上报 API返回值
// alibaba.alihealth.drug.code.error.report
//
// 提供码信息错误上报功能，用于数据校对
type AlibabaAlihealthDrugCodeErrorReportAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeErrorReportAPIResponseModel
}

// AlibabaAlihealthDrugCodeErrorReportAPIResponseModel is 码信息错误上报 成功返回结果
type AlibabaAlihealthDrugCodeErrorReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_error_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回结果code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 接口调用状态
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
