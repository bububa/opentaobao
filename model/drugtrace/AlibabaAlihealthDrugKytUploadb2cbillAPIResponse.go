package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytUploadb2cbillAPIResponse 快易通零售B2C API返回值
// alibaba.alihealth.drug.kyt.uploadb2cbill
//
// 快易通零售B2C单据上传
type AlibabaAlihealthDrugKytUploadb2cbillAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytUploadb2cbillAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUploadb2cbillAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytUploadb2cbillAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytUploadb2cbillAPIResponseModel is 快易通零售B2C 成功返回结果
type AlibabaAlihealthDrugKytUploadb2cbillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_uploadb2cbill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果值
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 调用结果
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 调用结果描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// success
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUploadb2cbillAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytUploadb2cbillAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytUploadb2cbillAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytUploadb2cbillAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytUploadb2cbillAPIResponse
func GetAlibabaAlihealthDrugKytUploadb2cbillAPIResponse() *AlibabaAlihealthDrugKytUploadb2cbillAPIResponse {
	return poolAlibabaAlihealthDrugKytUploadb2cbillAPIResponse.Get().(*AlibabaAlihealthDrugKytUploadb2cbillAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytUploadb2cbillAPIResponse 将 AlibabaAlihealthDrugKytUploadb2cbillAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytUploadb2cbillAPIResponse(v *AlibabaAlihealthDrugKytUploadb2cbillAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytUploadb2cbillAPIResponse.Put(v)
}
