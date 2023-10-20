package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerBillUploadAPIResponse 上传入出库单api API返回值
// alibaba.alihealth.tracecodeseller.bill.upload
//
// 上传入出库单api
type AlibabaAlihealthTracecodesellerBillUploadAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodesellerBillUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerBillUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTracecodesellerBillUploadAPIResponseModel).Reset()
}

// AlibabaAlihealthTracecodesellerBillUploadAPIResponseModel is 上传入出库单api 成功返回结果
type AlibabaAlihealthTracecodesellerBillUploadAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerBillUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = false
}

var poolAlibabaAlihealthTracecodesellerBillUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesellerBillUploadAPIResponse)
	},
}

// GetAlibabaAlihealthTracecodesellerBillUploadAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerBillUploadAPIResponse
func GetAlibabaAlihealthTracecodesellerBillUploadAPIResponse() *AlibabaAlihealthTracecodesellerBillUploadAPIResponse {
	return poolAlibabaAlihealthTracecodesellerBillUploadAPIResponse.Get().(*AlibabaAlihealthTracecodesellerBillUploadAPIResponse)
}

// ReleaseAlibabaAlihealthTracecodesellerBillUploadAPIResponse 将 AlibabaAlihealthTracecodesellerBillUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerBillUploadAPIResponse(v *AlibabaAlihealthTracecodesellerBillUploadAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerBillUploadAPIResponse.Put(v)
}
