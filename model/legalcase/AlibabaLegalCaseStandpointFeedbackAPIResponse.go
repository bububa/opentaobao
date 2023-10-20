package legalcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseStandpointFeedbackAPIResponse 新增或更新 反馈口径(采纳口径/不采纳口径) API返回值
// alibaba.legal.case.standpoint.feedback
//
// 新增或更新 反馈口径(采纳口径/不采纳口径)
type AlibabaLegalCaseStandpointFeedbackAPIResponse struct {
	model.CommonResponse
	AlibabaLegalCaseStandpointFeedbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalCaseStandpointFeedbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalCaseStandpointFeedbackAPIResponseModel).Reset()
}

// AlibabaLegalCaseStandpointFeedbackAPIResponseModel is 新增或更新 反馈口径(采纳口径/不采纳口径) 成功返回结果
type AlibabaLegalCaseStandpointFeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_standpoint_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误编码
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误信息描述
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 成功失败 true/false
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalCaseStandpointFeedbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errcode = ""
	m.Errmsg = ""
	m.IsSuccess = false
	m.Content = false
}

var poolAlibabaLegalCaseStandpointFeedbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalCaseStandpointFeedbackAPIResponse)
	},
}

// GetAlibabaLegalCaseStandpointFeedbackAPIResponse 从 sync.Pool 获取 AlibabaLegalCaseStandpointFeedbackAPIResponse
func GetAlibabaLegalCaseStandpointFeedbackAPIResponse() *AlibabaLegalCaseStandpointFeedbackAPIResponse {
	return poolAlibabaLegalCaseStandpointFeedbackAPIResponse.Get().(*AlibabaLegalCaseStandpointFeedbackAPIResponse)
}

// ReleaseAlibabaLegalCaseStandpointFeedbackAPIResponse 将 AlibabaLegalCaseStandpointFeedbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalCaseStandpointFeedbackAPIResponse(v *AlibabaLegalCaseStandpointFeedbackAPIResponse) {
	v.Reset()
	poolAlibabaLegalCaseStandpointFeedbackAPIResponse.Put(v)
}
