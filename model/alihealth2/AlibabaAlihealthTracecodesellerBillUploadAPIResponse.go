package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellerbilluploadAPIResponse 上传入出库单api API返回值
// alibaba.alihealth.tracecodeseller.bill.upload
//
// 上传入出库单api
type AlibabaalihealthtracecodesellerbilluploadAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthtracecodesellerbilluploadAPIResponseModel
}

// AlibabaalihealthtracecodesellerbilluploadAPIResponseModel is 上传入出库单api 成功返回结果
type AlibabaalihealthtracecodesellerbilluploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_bill_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
