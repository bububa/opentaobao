package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellerbillrootcodegetAPIResponse 获取最外层包装码 API返回值
// alibaba.alihealth.tracecodeseller.bill.rootcode.get
//
// 获取最外层包装码
type AlibabaalihealthtracecodesellerbillrootcodegetAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthtracecodesellerbillrootcodegetAPIResponseModel
}

// AlibabaalihealthtracecodesellerbillrootcodegetAPIResponseModel is 获取最外层包装码 成功返回结果
type AlibabaalihealthtracecodesellerbillrootcodegetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_bill_rootcode_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层码
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
