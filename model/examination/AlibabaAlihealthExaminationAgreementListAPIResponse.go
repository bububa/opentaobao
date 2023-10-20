package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationAgreementListAPIResponse isv协议获取 API返回值
// alibaba.alihealth.examination.agreement.list
//
// isv协议获取
type AlibabaAlihealthExaminationAgreementListAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationAgreementListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationAgreementListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationAgreementListAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationAgreementListAPIResponseModel is isv协议获取 成功返回结果
type AlibabaAlihealthExaminationAgreementListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_agreement_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果编码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 返回结果描述
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 返回的json格式数据
	Agreement *Agreement `json:"agreement,omitempty" xml:"agreement,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationAgreementListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResponseCode = ""
	m.ResponseMessage = ""
	m.Agreement = nil
}

var poolAlibabaAlihealthExaminationAgreementListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationAgreementListAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationAgreementListAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationAgreementListAPIResponse
func GetAlibabaAlihealthExaminationAgreementListAPIResponse() *AlibabaAlihealthExaminationAgreementListAPIResponse {
	return poolAlibabaAlihealthExaminationAgreementListAPIResponse.Get().(*AlibabaAlihealthExaminationAgreementListAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationAgreementListAPIResponse 将 AlibabaAlihealthExaminationAgreementListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationAgreementListAPIResponse(v *AlibabaAlihealthExaminationAgreementListAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationAgreementListAPIResponse.Put(v)
}
